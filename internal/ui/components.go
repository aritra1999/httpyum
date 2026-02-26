package ui

import (
	"fmt"
	"strings"

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
