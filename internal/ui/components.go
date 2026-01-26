package ui

import (
	"fmt"
	"regexp"
	"strings"

	"httpyum/internal/client"
	"httpyum/internal/parser"
)

func RenderRequestListItem(req parser.Request, isSelected bool, index int) string {
	var sb strings.Builder

	method := req.Method
	url := req.URL
	if len(url) > 60 {
		url = url[:57] + "..."
	}

	requestLine := fmt.Sprintf("%s %s", method, url)

	if isSelected {
		sb.WriteString(selectedStyle.Render(fmt.Sprintf("▶ %s", requestLine)))
	} else {
		sb.WriteString(normalStyle.Render(requestLine))
	}

	if req.Description != "" {
		sb.WriteString("\n")
		if isSelected {
			sb.WriteString(selectedStyle.Copy().PaddingLeft(4).Render(
				descriptionStyle.Render(fmt.Sprintf("  %s", req.Description)),
			))
		} else {
			sb.WriteString(normalStyle.Copy().PaddingLeft(6).Render(
				descriptionStyle.Render(req.Description),
			))
		}
	}

	return sb.String()
}

func RenderResponseBox(result *client.ExecutionResult, showHeaders bool, width int) string {
	return RenderResponseBoxWithVariables(result, showHeaders, nil, false, width)
}

func RenderResponseStaticSection(result *client.ExecutionResult, showHeaders bool, variables map[string]string, showVariables bool, width int) string {
	var sections []string

	if showVariables && variables != nil {
		reqSection := renderRequestSectionWithVariables(result, variables, width)
		sections = append(sections, reqSection)
	} else {
		reqSection := renderRequestSection(result)
		sections = append(sections, reqSection)
	}

	statusSection := renderStatusSection(result)
	sections = append(sections, statusSection)

	if showHeaders && result.Response != nil && len(result.Response.Headers) > 0 {
		headerSection := renderHeadersSection(result)
		sections = append(sections, headerSection)
	}

	paddedSections := make([]string, len(sections))
	for i, section := range sections {
		lines := strings.Split(section, "\n")
		for j, line := range lines {
			lines[j] = "  " + line
		}
		paddedSections[i] = strings.Join(lines, "\n")
	}

	separator := "\n" + mutedStyle.Render(strings.Repeat("─", width-4)) + "\n"
	content := strings.Join(paddedSections, separator)

	return boxStyle.Width(width - 4).Render(content)
}

func RenderResponseBodyContent(result *client.ExecutionResult) string {
	bodySection := renderBodySection(result)

	lines := strings.Split(bodySection, "\n")
	for j, line := range lines {
		lines[j] = "  " + line
	}

	return strings.Join(lines, "\n")
}

func RenderResponseBoxWithVariables(result *client.ExecutionResult, showHeaders bool, variables map[string]string, showVariables bool, width int) string {
	var sections []string

	if showVariables && variables != nil {
		reqSection := renderRequestSectionWithVariables(result, variables, width)
		sections = append(sections, reqSection)
	} else {
		reqSection := renderRequestSection(result)
		sections = append(sections, reqSection)
	}

	statusSection := renderStatusSection(result)
	sections = append(sections, statusSection)

	if showHeaders && result.Response != nil && len(result.Response.Headers) > 0 {
		headerSection := renderHeadersSection(result)
		sections = append(sections, headerSection)
	}

	bodySection := renderBodySection(result)
	sections = append(sections, bodySection)

	paddedSections := make([]string, len(sections))
	for i, section := range sections {
		lines := strings.Split(section, "\n")
		for j, line := range lines {
			lines[j] = "  " + line
		}
		paddedSections[i] = strings.Join(lines, "\n")
	}

	separator := "\n" + mutedStyle.Render(strings.Repeat("─", width-4)) + "\n"
	content := strings.Join(paddedSections, separator)

	return boxStyle.Width(width - 4).Render(content)
}

func renderRequestSection(result *client.ExecutionResult) string {
	var sb strings.Builder

	sb.WriteString(infoStyle.Bold(true).Render("Request"))
	sb.WriteString("\n")
	sb.WriteString(successStyle.Render(fmt.Sprintf("%s %s", result.Request.Method, result.Request.URL)))

	if len(result.Request.Headers) > 0 {
		sb.WriteString("\n")
		for key, value := range result.Request.Headers {
			if len(value) > 80 {
				value = value[:77] + "..."
			}
			sb.WriteString(fmt.Sprintf("%s: %s\n", headerKeyStyle.Render(key), headerValueStyle.Render(value)))
		}
	}

	return sb.String()
}

func renderRequestSectionWithVariables(result *client.ExecutionResult, allVariables map[string]string, totalWidth int) string {
	contentWidth := totalWidth - 4
	leftWidth := int(float64(contentWidth) * 0.65)
	rightWidth := contentWidth - leftWidth - 3

	var leftSb strings.Builder
	leftSb.WriteString(infoStyle.Bold(true).Render("Request"))
	leftSb.WriteString("\n")
	leftSb.WriteString(successStyle.Render(fmt.Sprintf("%s %s", result.Request.Method, result.Request.URL)))

	if len(result.Request.Headers) > 0 {
		leftSb.WriteString("\n")
		for key, value := range result.Request.Headers {
			maxValueLen := leftWidth - len(key) - 4
			if maxValueLen > 0 && len(value) > maxValueLen {
				value = value[:maxValueLen-3] + "..."
			}
			leftSb.WriteString(fmt.Sprintf("%s: %s\n", headerKeyStyle.Render(key), headerValueStyle.Render(value)))
		}
	}

	var rightSb strings.Builder
	rightSb.WriteString(infoStyle.Bold(true).Render("Variables Used"))
	rightSb.WriteString("\n")

	usedVars := parser.ExtractUsedVariables(result.Request, allVariables)

	if len(usedVars) == 0 {
		rightSb.WriteString(mutedStyle.Render("(none)"))
	} else {
		for key, value := range usedVars {
			maskedValue := maskValue(value)
			displayKey := key
			if strings.HasPrefix(key, "$dotenv_") {
				displayKey = "$" + strings.TrimPrefix(key, "$dotenv_")
			}
			line := fmt.Sprintf("%s = %s", displayKey, maskedValue)
			if len(line) > rightWidth {
				line = line[:rightWidth-3] + "..."
			}
			rightSb.WriteString(fmt.Sprintf("%s\n", mutedStyle.Render(line)))
		}
	}

	leftLines := strings.Split(strings.TrimSuffix(leftSb.String(), "\n"), "\n")
	rightLines := strings.Split(strings.TrimSuffix(rightSb.String(), "\n"), "\n")

	maxLines := len(leftLines)
	if len(rightLines) > maxLines {
		maxLines = len(rightLines)
	}

	var output strings.Builder
	divider := mutedStyle.Render("│")

	for i := 0; i < maxLines; i++ {
		leftLine := ""
		if i < len(leftLines) {
			leftLine = leftLines[i]
		}

		visualLen := visualLength(leftLine)
		output.WriteString(leftLine)

		if visualLen < leftWidth {
			output.WriteString(strings.Repeat(" ", leftWidth-visualLen))
		}

		output.WriteString(" " + divider + " ")

		if i < len(rightLines) {
			output.WriteString(rightLines[i])
		}

		output.WriteString("\n")
	}

	return strings.TrimSuffix(output.String(), "\n")
}

func visualLength(s string) int {
	ansiRegex := regexp.MustCompile(`\x1b\[[0-9;]*m`)
	cleaned := ansiRegex.ReplaceAllString(s, "")
	return len(cleaned)
}

func renderStatusSection(result *client.ExecutionResult) string {
	var sb strings.Builder

	sb.WriteString(infoStyle.Bold(true).Render("Response"))
	sb.WriteString("\n")

	if result.Error != nil {
		sb.WriteString(errorStyle.Render(fmt.Sprintf("Error: %v", result.Error)))
		if result.Response != nil {
			sb.WriteString(fmt.Sprintf(" | %s", mutedStyle.Render(result.Response.Duration.String())))
		}
	} else if result.Response != nil {
		statusStyle := StatusCodeStyle(result.Response.StatusCode, result.Response.Status)
		sb.WriteString(statusStyle.Render(result.Response.Status))
		sb.WriteString(fmt.Sprintf(" | %s", mutedStyle.Render(result.Response.Duration.String())))
		sb.WriteString(fmt.Sprintf(" | %s", mutedStyle.Render(client.FormatSize(result.Response.Size))))
	}

	return sb.String()
}

func renderHeadersSection(result *client.ExecutionResult) string {
	var sb strings.Builder

	sb.WriteString(infoStyle.Bold(true).Render("Response Headers"))
	sb.WriteString(mutedStyle.Render(" (press 'h' to hide)"))
	sb.WriteString("\n")

	for key, values := range result.Response.Headers {
		value := strings.Join(values, ", ")
		if len(value) > 80 {
			value = value[:77] + "..."
		}
		sb.WriteString(fmt.Sprintf("%s: %s\n", headerKeyStyle.Render(key), headerValueStyle.Render(value)))
	}

	return strings.TrimSuffix(sb.String(), "\n")
}

func renderBodySection(result *client.ExecutionResult) string {
	var sb strings.Builder

	sb.WriteString(infoStyle.Bold(true).Render("Response Body"))

	if result.Response != nil && client.IsJSON(result.Response.ContentType) {
		sb.WriteString(mutedStyle.Render(" (press 'f' to explore interactively)"))
	}
	sb.WriteString("\n")

	if result.Response == nil || len(result.Response.Body) == 0 {
		sb.WriteString(mutedStyle.Render("(empty)"))
	} else {
		body := string(result.Response.Body)

		if client.IsJSON(result.Response.ContentType) {
			prettyJSON, err := client.PrettyPrintJSON(result.Response.Body)
			if err == nil {
				body = prettyJSON
			}
		}

		maxBodyLength := 2000
		if len(body) > maxBodyLength {
			body = body[:maxBodyLength] + "\n" + mutedStyle.Render("... (truncated)")
		}

		sb.WriteString(body)
	}

	return sb.String()
}

func RenderHelpBar(currentView ViewType) string {
	var shortcuts []string

	switch currentView {
	case ViewList:
		shortcuts = []string{
			"↑/↓: navigate",
			"/: filter",
			"enter: execute",
			"q: quit",
		}
	case ViewResponse:
		shortcuts = []string{
			"↑/↓: scroll",
			"f: interactive JSON",
			"h: toggle headers",
			"v: toggle variables",
			"esc/b: back",
			"q: quit",
		}
	case ViewLoading:
		shortcuts = []string{
			"Loading...",
		}
	case ViewError:
		shortcuts = []string{
			"esc/b: back to list",
			"q: quit",
		}
	}

	return helpStyle.Render(strings.Join(shortcuts, " • "))
}

func RenderSpinner(frame int) string {
	spinners := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	spinner := spinners[frame%len(spinners)]
	return infoStyle.Render(fmt.Sprintf("%s Loading...", spinner))
}

func maskValue(value string) string {
	if len(value) <= 3 {
		return "***"
	}
	return "..." + value[len(value)-3:]
}
