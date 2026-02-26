package ui

import (
	"fmt"
	"sort"
	"strings"

	"httpyum/internal/client"
	"httpyum/internal/parser"

	"github.com/charmbracelet/lipgloss"
)

// RenderOpts controls what sections appear in the response content.
type RenderOpts struct {
	ShowHeaders    bool
	ShowVariables  bool
	Variables      map[string]string
	ContentWidth   int
	ViewportHeight int
}

var borderStyle = lipgloss.NewStyle().Foreground(colorPrimary)

// RenderResponseContent produces the viewport content with box side borders,
// proper junction characters where column dividers meet separators, and
// padding lines that preserve the column divider to the bottom of the box.
func RenderResponseContent(result *client.ExecutionResult, opts RenderOpts) string {
	cw := opts.ContentWidth
	innerWidth := cw + 2 // 1 space padding each side

	// Column layout (shared by headers and body two-column sections)
	leftWidth := cw / 2
	rightWidth := cw - leftWidth - 3
	if rightWidth < 0 {
		rightWidth = 0
	}

	side := borderStyle.Render("│")
	colDivider := mutedStyle.Render("│")

	// Separator builders
	plainSep := borderStyle.Render("├" + strings.Repeat("─", innerWidth) + "┤")
	colSep := func(junction string) string {
		leftDashes := leftWidth + 2   // space + left content + space before divider
		rightDashes := rightWidth + 2 // space after divider + right content + space
		return borderStyle.Render("├" + strings.Repeat("─", leftDashes) + junction + strings.Repeat("─", rightDashes) + "┤")
	}

	wrapLine := func(line string) string {
		vl := visualLength(line)
		pad := ""
		if vl < cw {
			pad = strings.Repeat(" ", cw-vl)
		}
		return side + " " + line + pad + " " + side
	}

	wrapSection := func(section string) []string {
		lines := strings.Split(section, "\n")
		wrapped := make([]string, len(lines))
		for i, line := range lines {
			wrapped[i] = wrapLine(line)
		}
		return wrapped
	}

	// Determine which sections exist
	hasReqHeaders := len(result.Request.Headers) > 0
	hasResHeaders := opts.ShowHeaders && result.Response != nil && len(result.Response.Headers) > 0
	hasHeaders := hasReqHeaders || hasResHeaders
	hasReqBody := result.Request.Body != ""

	// Build output lines
	var allLines []string

	// Section 1: Request details (single column)
	allLines = append(allLines, wrapSection(renderRequestDetails(result, opts))...)

	// Section 2: Headers (two-column)
	if hasHeaders {
		allLines = append(allLines, colSep("┬"))
		allLines = append(allLines, wrapSection(renderHeadersTwoColumn(result, hasResHeaders, cw))...)
	}

	// Section 3: Body
	if hasReqBody {
		if hasHeaders {
			allLines = append(allLines, colSep("┼"))
		} else {
			allLines = append(allLines, colSep("┬"))
		}
		allLines = append(allLines, wrapSection(renderBodyTwoColumn(result, cw))...)
	} else {
		if hasHeaders {
			allLines = append(allLines, colSep("┴"))
		} else {
			allLines = append(allLines, plainSep)
		}
		allLines = append(allLines, wrapSection(renderBody(result))...)
	}

	// Determine if last section is two-column (for padding)
	lastTwoCol := hasReqBody

	// Pad to viewport height, preserving column divider if last section is two-col
	if opts.ViewportHeight > 0 {
		var padLine string
		if lastTwoCol {
			padLine = side + " " + strings.Repeat(" ", leftWidth) + " " + colDivider + " " + strings.Repeat(" ", rightWidth) + " " + side
		} else {
			padLine = side + strings.Repeat(" ", innerWidth) + side
		}
		for len(allLines) < opts.ViewportHeight {
			allLines = append(allLines, padLine)
		}
	}

	return strings.Join(allLines, "\n")
}

// RenderTopBorder renders ╭───────────╮
func RenderTopBorder(width int) string {
	innerWidth := max(width-2, 0)
	return borderStyle.Render("╭" + strings.Repeat("─", innerWidth) + "╮")
}

// RenderBottomBorder renders ╰──┴──── 200 OK | 143ms | 2.1 KB ────╯
// colPos is the junction position for ┴ (0 means no junction).
func RenderBottomBorder(result *client.ExecutionResult, width, colPos int) string {
	innerWidth := max(width-2, 0)

	label := ""
	if result.Error != nil {
		label = " Error "
	} else if result.Response != nil {
		label = fmt.Sprintf(" %s | %s | %s ",
			result.Response.Status,
			result.Response.Duration.String(),
			client.FormatSize(result.Response.Size),
		)
	}

	// Build the inner dashes, inserting ┴ junction if needed
	buildDashes := func(n int) string {
		if colPos <= 0 || colPos >= n {
			return strings.Repeat("─", n)
		}
		return strings.Repeat("─", colPos) + "┴" + strings.Repeat("─", n-colPos-1)
	}

	if label == "" {
		return borderStyle.Render("╰" + buildDashes(innerWidth) + "╯")
	}

	var labelStyled string
	if result.Error != nil {
		labelStyled = errorStyle.Render(label)
	} else {
		statusStyle := StatusCodeStyle(result.Response.StatusCode, result.Response.Status)
		labelStyled = statusStyle.Render(label)
	}

	labelVisual := visualLength(labelStyled)
	trailDashes := 3
	leadDashes := innerWidth - labelVisual - trailDashes
	if leadDashes < 1 {
		leadDashes = 1
	}

	return borderStyle.Render("╰"+buildDashes(leadDashes)) + labelStyled + borderStyle.Render(strings.Repeat("─", trailDashes)+"╯")
}

func twoColumn(leftLines, rightLines []string, leftWidth, rightWidth int) string {
	maxLines := len(leftLines)
	if len(rightLines) > maxLines {
		maxLines = len(rightLines)
	}

	var sb strings.Builder
	divider := mutedStyle.Render("│")

	for i := 0; i < maxLines; i++ {
		leftLine := ""
		if i < len(leftLines) {
			leftLine = leftLines[i]
		}
		rightLine := ""
		if i < len(rightLines) {
			rightLine = rightLines[i]
		}

		lv := visualLength(leftLine)
		sb.WriteString(leftLine)
		if lv < leftWidth {
			sb.WriteString(strings.Repeat(" ", leftWidth-lv))
		}

		sb.WriteString(" " + divider + " ")

		rv := visualLength(rightLine)
		sb.WriteString(rightLine)
		if rv < rightWidth {
			sb.WriteString(strings.Repeat(" ", rightWidth-rv))
		}

		sb.WriteString("\n")
	}

	return strings.TrimSuffix(sb.String(), "\n")
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	if maxLen <= 3 {
		return s[:maxLen]
	}
	return s[:maxLen-3] + "..."
}

func wrapText(text string, width int) []string {
	if width <= 0 {
		return []string{text}
	}
	lines := strings.Split(text, "\n")
	var result []string
	for _, line := range lines {
		for len(line) > width {
			result = append(result, line[:width])
			line = line[width:]
		}
		result = append(result, line)
	}
	return result
}

func visualLength(s string) int {
	return lipgloss.Width(s)
}

func maskValue(value string) string {
	if len(value) <= 3 {
		return "***"
	}
	return "..." + value[len(value)-3:]
}

// --- Section renderers ---

// renderRequestDetails renders method, URL, and variables.
func renderRequestDetails(result *client.ExecutionResult, opts RenderOpts) string {
	var sb strings.Builder

	sb.WriteString(sectionTitleStyle.Render("Request"))
	sb.WriteString("\n")
	sb.WriteString(successStyle.Render(fmt.Sprintf("%s %s", result.Request.Method, result.Request.URL)))

	if opts.ShowVariables && opts.Variables != nil {
		sb.WriteString("\n\n")
		sb.WriteString(buildVariablesText(result, opts.Variables, opts.ContentWidth))
	}

	return sb.String()
}

// renderHeadersTwoColumn renders request headers (left) and response headers (right).
func renderHeadersTwoColumn(result *client.ExecutionResult, showResHeaders bool, totalWidth int) string {
	leftWidth := totalWidth / 2
	rightWidth := totalWidth - leftWidth - 3

	// Left: Request Headers
	var leftSb strings.Builder
	leftSb.WriteString(sectionTitleStyle.Render("Request Headers"))
	for _, h := range result.Request.Headers {
		leftSb.WriteString("\n")
		value := h.Value
		maxValueLen := leftWidth - len(h.Key) - 4
		if maxValueLen > 0 && len(value) > maxValueLen {
			value = truncate(value, maxValueLen)
		}
		leftSb.WriteString(fmt.Sprintf("%s: %s", headerKeyStyle.Render(h.Key), headerValueStyle.Render(value)))
	}

	// Right: Response Headers
	var rightSb strings.Builder
	if showResHeaders {
		rightSb.WriteString(sectionTitleStyle.Render("Response Headers"))
		rightSb.WriteString(mutedStyle.Render(" ('h' to hide)"))

		keys := make([]string, 0, len(result.Response.Headers))
		for key := range result.Response.Headers {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		for _, key := range keys {
			rightSb.WriteString("\n")
			value := strings.Join(result.Response.Headers[key], ", ")
			if rightWidth > 0 && len(value) > rightWidth {
				value = truncate(value, rightWidth)
			}
			rightSb.WriteString(fmt.Sprintf("%s: %s", headerKeyStyle.Render(key), headerValueStyle.Render(value)))
		}
	}

	leftLines := strings.Split(leftSb.String(), "\n")
	rightLines := strings.Split(rightSb.String(), "\n")

	return twoColumn(leftLines, rightLines, leftWidth, rightWidth)
}

func buildVariablesText(result *client.ExecutionResult, allVariables map[string]string, maxWidth int) string {
	var sb strings.Builder
	sb.WriteString(sectionTitleStyle.Render("Variables Used"))

	usedVars := parser.ExtractUsedVariables(result.Request, allVariables)

	if len(usedVars) == 0 {
		sb.WriteString("\n")
		sb.WriteString(mutedStyle.Render("(none)"))
	} else {
		varKeys := make([]string, 0, len(usedVars))
		for key := range usedVars {
			varKeys = append(varKeys, key)
		}
		sort.Strings(varKeys)
		for _, key := range varKeys {
			value := usedVars[key]
			sb.WriteString("\n")
			maskedValue := maskValue(value)
			displayKey := key
			if strings.HasPrefix(key, "$dotenv_") {
				displayKey = "$" + strings.TrimPrefix(key, "$dotenv_")
			}
			line := fmt.Sprintf("%s = %s", displayKey, maskedValue)
			if maxWidth > 0 && len(line) > maxWidth {
				line = truncate(line, maxWidth)
			}
			sb.WriteString(mutedStyle.Render(line))
		}
	}

	return sb.String()
}

func renderBody(result *client.ExecutionResult) string {
	var sb strings.Builder

	sb.WriteString(sectionTitleStyle.Render("Response Body"))

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

		sb.WriteString(body)
	}

	return sb.String()
}

func renderBodyTwoColumn(result *client.ExecutionResult, totalWidth int) string {
	leftWidth := totalWidth / 2
	rightWidth := totalWidth - leftWidth - 3

	// Left column: Request Body
	var leftSb strings.Builder
	leftSb.WriteString(sectionTitleStyle.Render("Request Body"))
	leftSb.WriteString("\n")

	if result.Request.Body == "" {
		leftSb.WriteString(mutedStyle.Render("(empty)"))
	} else {
		body := result.Request.Body

		if len(body) > 0 && (body[0] == '{' || body[0] == '[') {
			prettyJSON, err := client.PrettyPrintJSON([]byte(body))
			if err == nil {
				body = prettyJSON
			}
		}

		wrapped := wrapText(body, leftWidth)
		leftSb.WriteString(strings.Join(wrapped, "\n"))
	}

	// Right column: Response Body
	var rightSb strings.Builder
	rightSb.WriteString(sectionTitleStyle.Render("Response Body"))

	if result.Response != nil && client.IsJSON(result.Response.ContentType) {
		rightSb.WriteString(mutedStyle.Render(" ('f' to explore)"))
	}
	rightSb.WriteString("\n")

	if result.Response == nil || len(result.Response.Body) == 0 {
		rightSb.WriteString(mutedStyle.Render("(empty)"))
	} else {
		body := string(result.Response.Body)

		if client.IsJSON(result.Response.ContentType) {
			prettyJSON, err := client.PrettyPrintJSON(result.Response.Body)
			if err == nil {
				body = prettyJSON
			}
		}

		wrapped := wrapText(body, rightWidth)
		rightSb.WriteString(strings.Join(wrapped, "\n"))
	}

	leftLines := strings.Split(leftSb.String(), "\n")
	rightLines := strings.Split(rightSb.String(), "\n")

	return twoColumn(leftLines, rightLines, leftWidth, rightWidth)
}
