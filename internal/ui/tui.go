package ui

import (
	"os"
	"os/exec"
	"strings"
	"time"

	"httpyum/internal/client"
	"httpyum/internal/parser"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type ViewType string

const (
	ViewList     ViewType = "list"
	ViewResponse ViewType = "response"
	ViewLoading  ViewType = "loading"
	ViewError    ViewType = "error"
)

type requestItem struct {
	request parser.Request
}

func (i requestItem) FilterValue() string {
	return i.request.Method + " " + i.request.URL
}

func (i requestItem) Title() string {
	return i.request.Method + " " + i.request.URL
}

func (i requestItem) Description() string {
	return i.request.Description
}

type Model struct {
	ParsedFile          *parser.ParsedFile
	Requests            []parser.Request
	Variables           map[string]string
	list                list.Model
	viewport            viewport.Model
	CurrentView         ViewType
	LastResult          *client.ExecutionResult
	ShowHeaders         bool
	ShowVariables       bool
	ErrorMsg            string
	Width               int
	Height              int
	SpinnerFrame        int
	executor            *client.Executor
	cachedStaticSection string
}

func NewModel(parsedFile *parser.ParsedFile, envVars map[string]string, showHeaders bool) Model {
	variables := parser.BuildVariableMap(parsedFile.Variables, envVars)

	items := make([]list.Item, len(parsedFile.Requests))
	for i, req := range parsedFile.Requests {
		items[i] = requestItem{request: req}
	}

	delegate := itemDelegate{}
	requestList := list.New(items, delegate, 0, listHeight)
	requestList.Title = ""
	requestList.SetShowStatusBar(true)
	requestList.SetFilteringEnabled(true)
	requestList.SetShowHelp(true)
	requestList.DisableQuitKeybindings()

	requestList.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{}
	}
	requestList.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{}
	}

	vp := viewport.New(80, 20)
	vp.Style = viewportStyle

	return Model{
		ParsedFile:    parsedFile,
		Requests:      parsedFile.Requests,
		Variables:     variables,
		list:          requestList,
		viewport:      vp,
		CurrentView:   ViewList,
		ShowHeaders:   showHeaders,
		ShowVariables: true,
		Width:         80,
		Height:        24,
		SpinnerFrame:  0,
		executor:      client.NewExecutor(variables),
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

type executeFinishedMsg struct {
	result *client.ExecutionResult
}

type tickMsg time.Time

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.CurrentView == ViewList {
			switch msg.String() {
			case "q", "ctrl+c":
				return m, tea.Quit
			case "enter":
				if selectedItem, ok := m.list.SelectedItem().(requestItem); ok {
					m.CurrentView = ViewLoading
					m.SpinnerFrame = 0
					return m, tea.Batch(executeRequest(m.executor, &selectedItem.request), tick())
				}
			default:
				m.list, cmd = m.list.Update(msg)
				return m, cmd
			}
		} else {
			return m.handleKeyPress(msg)
		}

	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

		if m.CurrentView == ViewList {
			h, v := docStyle.GetFrameSize()
			m.list.SetSize(msg.Width-h, msg.Height-v)
		}

		m.viewport.Width = msg.Width - 8
		if m.CurrentView == ViewResponse {
			if m.LastResult != nil {
				m.cachedStaticSection = RenderResponseStaticSection(m.LastResult, m.ShowHeaders, m.Variables, m.ShowVariables, m.Width)
			}
			m.updateViewportHeight()
		}
		return m, nil

	case executeFinishedMsg:
		m.LastResult = msg.result
		m.CurrentView = ViewResponse

		m.list.ResetFilter()

		m.updateViewportContent()
		return m, nil

	case tickMsg:
		m.SpinnerFrame++
		if m.CurrentView == ViewLoading {
			return m, tick()
		}
		return m, nil

	case jsonViewerErrorMsg:
		m.ErrorMsg = "Error opening JSON viewer: " + msg.err.Error()
		m.CurrentView = ViewResponse
		return m, nil
	}

	return m, nil
}

func (m Model) handleKeyPress(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch m.CurrentView {
	case ViewResponse:
		return m.handleResponseKeys(msg)
	case ViewError:
		return m.handleErrorKeys(msg)
	default:
		return m, nil
	}
}

func (m Model) handleResponseKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit

	case "h":
		m.ShowHeaders = !m.ShowHeaders
		m.updateViewportContent()
		return m, nil

	case "v":
		m.ShowVariables = !m.ShowVariables
		m.updateViewportContent()
		return m, nil

	case "b", "esc":
		m.CurrentView = ViewList
		return m, nil

	case "f":
		if m.LastResult != nil && m.LastResult.Response != nil {
			if client.IsJSON(m.LastResult.Response.ContentType) {
				if _, err := exec.LookPath("jless"); err == nil {
					return m, openInJSONViewer(m.LastResult.Response.Body, "jless")
				} else if _, err := exec.LookPath("fx"); err == nil {
					return m, openInJSONViewer(m.LastResult.Response.Body, "fx")
				} else {
					m.ErrorMsg = "No JSON viewer installed. Install jless (recommended) or fx:\n  brew install jless\n  brew install fx"
					m.CurrentView = ViewError
					return m, nil
				}
			}
		}
		return m, nil

	case "up", "down", "pgup", "pgdown", "home", "end", "k", "j":
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) handleErrorKeys(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q", "ctrl+c":
		return m, tea.Quit

	case "b", "esc":
		m.CurrentView = ViewList
	}

	return m, nil
}

func (m Model) View() string {
	switch m.CurrentView {
	case ViewList:
		return m.RenderListView()
	case ViewResponse:
		return m.RenderResponseView()
	case ViewLoading:
		return m.RenderLoadingView()
	case ViewError:
		return m.RenderErrorView()
	default:
		return "Unknown view"
	}
}

func executeRequest(executor *client.Executor, req *parser.Request) tea.Cmd {
	return func() tea.Msg {
		result := executor.Execute(req)
		return executeFinishedMsg{result: result}
	}
}

func tick() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func openInJSONViewer(jsonData []byte, viewer string) tea.Cmd {
	tmpfile, err := os.CreateTemp("", "httpyum-*.json")
	if err != nil {
		return func() tea.Msg {
			return jsonViewerErrorMsg{err: err}
		}
	}

	if _, err := tmpfile.Write(jsonData); err != nil {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
		return func() tea.Msg {
			return jsonViewerErrorMsg{err: err}
		}
	}
	tmpfile.Close()

	tmpFileName := tmpfile.Name()

	c := exec.Command(viewer, tmpFileName)
	return tea.ExecProcess(c, func(err error) tea.Msg {
		os.Remove(tmpFileName)
		return nil
	})
}

type jsonViewerErrorMsg struct {
	err error
}

func (m *Model) updateViewportContent() {
	if m.LastResult == nil {
		return
	}

	m.cachedStaticSection = RenderResponseStaticSection(m.LastResult, m.ShowHeaders, m.Variables, m.ShowVariables, m.Width)

	bodyContent := RenderResponseBodyContent(m.LastResult)
	m.viewport.SetContent(bodyContent)

	m.updateViewportHeight()
}

func (m *Model) updateViewportHeight() {
	if m.LastResult == nil {
		return
	}

	staticHeight := strings.Count(m.cachedStaticSection, "\n") + 1

	reservedSpace := staticHeight + 8
	availableHeight := m.Height - reservedSpace

	if availableHeight < 5 {
		availableHeight = 5
	}

	m.viewport.Height = availableHeight
}
