package ui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const listHeight = 14

var (
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("#9C59D1"))
	dimmedItemStyle   = lipgloss.NewStyle().PaddingLeft(4).Foreground(colorMuted)
)

type itemDelegate struct{}

func (d itemDelegate) Height() int { return 2 }

func (d itemDelegate) Spacing() int { return 1 }

func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(requestItem)
	if !ok {
		return
	}

	method := i.request.Method
	var methodStyle lipgloss.Style
	switch method {
	case "GET":
		methodStyle = lipgloss.NewStyle().Foreground(colorSecondary).Bold(true)
	case "POST":
		methodStyle = lipgloss.NewStyle().Foreground(colorAccent).Bold(true)
	case "PUT":
		methodStyle = lipgloss.NewStyle().Foreground(colorWarning).Bold(true)
	case "DELETE":
		methodStyle = lipgloss.NewStyle().Foreground(colorError).Bold(true)
	case "PATCH":
		methodStyle = lipgloss.NewStyle().Foreground(colorPrimary).Bold(true)
	default:
		methodStyle = lipgloss.NewStyle().Foreground(colorText).Bold(true)
	}

	url := i.request.URL
	maxWidth := m.Width() - 20
	if maxWidth > 10 && len(url) > maxWidth {
		url = url[:maxWidth-3] + "..."
	} else if maxWidth <= 10 && len(url) > 30 {
		url = url[:27] + "..."
	}

	if index == m.Index() {
		line := selectedItemStyle.Render(fmt.Sprintf("│ %s %s",
			methodStyle.Render(method),
			url))
		fmt.Fprint(w, line)

		if i.request.Description != "" {
			desc := descriptionStyle.Render(i.request.Description)
			fmt.Fprintf(w, "%s", "\n"+selectedItemStyle.Render("│ ")+desc)
		}
	} else {
		line := itemStyle.Render(fmt.Sprintf("%s %s",
			methodStyle.Render(method),
			mutedStyle.Render(url)))
		fmt.Fprint(w, line)

		if i.request.Description != "" {
			desc := descriptionStyle.Render(i.request.Description)
			fmt.Fprintf(w, "%s", "\n"+dimmedItemStyle.Render(desc))
		}
	}
}
