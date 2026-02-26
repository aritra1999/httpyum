package ui

import (
	"fmt"
	"strings"
)

func (m Model) RenderListView() string {
	if len(m.Requests) == 0 {
		return titleStyle.Render("httpyum - HTTP Request Runner") + "\n\n" +
			mutedStyle.Render("No requests found in the file.")
	}

	return docStyle.Render(m.list.View())
}

func (m Model) RenderResponseView() string {
	if m.LastResult == nil {
		return errorStyle.Render("No response to display")
	}

	bw := m.boxWidth()
	margin := " " // 1-space left margin

	// Column junction position for bottom border (0 = no junction)
	colPos := 0
	if m.LastResult.Request.Body != "" {
		cw := m.contentWidth()
		leftWidth := cw / 2
		colPos = leftWidth + 2 // padding + left content + space before divider
	}

	// Add margin to each viewport line
	vpLines := strings.Split(m.viewport.View(), "\n")
	for i, line := range vpLines {
		vpLines[i] = margin + line
	}

	var sb strings.Builder
	sb.WriteString(margin + RenderTopBorder(bw))
	sb.WriteString("\n")
	sb.WriteString(strings.Join(vpLines, "\n"))
	sb.WriteString("\n")
	sb.WriteString(margin + RenderBottomBorder(m.LastResult, bw, colPos))
	sb.WriteString("\n")
	sb.WriteString(RenderHelpBar(ViewResponse))

	return sb.String()
}

func (m Model) RenderLoadingView() string {
	var sb strings.Builder

	sb.WriteString(RenderSpinner(m.SpinnerFrame))
	sb.WriteString("\n\n")

	if selectedItem, ok := m.list.SelectedItem().(requestItem); ok {
		sb.WriteString(mutedStyle.Render(fmt.Sprintf("Executing: %s %s", selectedItem.request.Method, selectedItem.request.URL)))
	}

	return docStyle.Render(sb.String())
}

func (m Model) RenderErrorView() string {
	var sb strings.Builder

	sb.WriteString(errorStyle.Render("Error: "))
	sb.WriteString(m.ErrorMsg)

	sb.WriteString("\n\n")
	sb.WriteString(RenderHelpBar(ViewError))

	return sb.String()
}
