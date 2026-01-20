package ui

import (
	"fmt"
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
	var sections []string

	reqSection := renderRequestSection(result)
	sections = append(sections, reqSection)

	statusSection := renderStatusSection(result)
	sections = append(sections, statusSection)

	if showHeaders && result.Response != nil && len(result.Response.Headers) > 0 {
		headerSection := renderHeadersSection(result)
		sections = append(sections, headerSection)
	}

	bodySection := renderBodySection(result)
	sections = append(sections, bodySection)

	separator := "\n" + mutedStyle.Render(strings.Repeat("─", width-8)) + "\n"
	content := strings.Join(sections, separator)

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
			"f: interactive JSON",
			"h: toggle headers",
			"esc/b: back to list",
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
