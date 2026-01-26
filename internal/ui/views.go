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

	var sb strings.Builder

	sb.WriteString(m.cachedStaticSection)

	sb.WriteString("\n")
	sb.WriteString(boxStyle.Width(m.Width - 4).Render(m.viewport.View()))

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
