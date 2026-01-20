package ui

import "github.com/charmbracelet/lipgloss"

var (
	colorPrimary   = lipgloss.Color("#7C3AED")
	colorSecondary = lipgloss.Color("#10B981")
	colorAccent    = lipgloss.Color("#3B82F6")
	colorWarning   = lipgloss.Color("#F59E0B")
	colorError     = lipgloss.Color("#EF4444")
	colorMuted     = lipgloss.Color("#6B7280")
	colorText      = lipgloss.Color("#F3F4F6")

	colorStatus2xx = lipgloss.Color("#10B981")
	colorStatus3xx = lipgloss.Color("#3B82F6")
	colorStatus4xx = lipgloss.Color("#F59E0B")
	colorStatus5xx = lipgloss.Color("#EF4444")
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorPrimary).
			MarginBottom(1)

	selectedStyle = lipgloss.NewStyle().
			Foreground(colorSecondary).
			Bold(true).
			PaddingLeft(2)

	normalStyle = lipgloss.NewStyle().
			Foreground(colorText).
			PaddingLeft(4)

	descriptionStyle = lipgloss.NewStyle().
				Foreground(colorMuted).
				Italic(true)

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorPrimary).
			Padding(1, 2)

	headerKeyStyle = lipgloss.NewStyle().
			Foreground(colorAccent).
			Bold(true)

	headerValueStyle = lipgloss.NewStyle().
				Foreground(colorText)

	errorStyle = lipgloss.NewStyle().
			Foreground(colorError).
			Bold(true)

	successStyle = lipgloss.NewStyle().
			Foreground(colorSecondary).
			Bold(true)

	infoStyle = lipgloss.NewStyle().
			Foreground(colorAccent)

	mutedStyle = lipgloss.NewStyle().
			Foreground(colorMuted)

	helpStyle = lipgloss.NewStyle().
			Foreground(colorMuted).
			MarginTop(1)

	docStyle = lipgloss.NewStyle().
			Margin(1, 2)
)

func StatusCodeColor(statusCode int) lipgloss.Color {
	switch {
	case statusCode >= 200 && statusCode < 300:
		return colorStatus2xx
	case statusCode >= 300 && statusCode < 400:
		return colorStatus3xx
	case statusCode >= 400 && statusCode < 500:
		return colorStatus4xx
	case statusCode >= 500:
		return colorStatus5xx
	default:
		return colorMuted
	}
}

func StatusCodeStyle(statusCode int, status string) lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(StatusCodeColor(statusCode)).
		Bold(true)
}
