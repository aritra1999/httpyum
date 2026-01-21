package main

import (
	"fmt"
	"os"

	"httpyum/internal/config"
	"httpyum/internal/parser"
	"httpyum/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	file, err := os.Open(cfg.FilePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	parsedFile, err := parser.Parse(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file: %v\n", err)
		os.Exit(1)
	}

	if len(parsedFile.Requests) == 0 {
		fmt.Fprintf(os.Stderr, "No HTTP requests found in file: %s\n", cfg.FilePath)
		os.Exit(1)
	}

	envVars := parser.LoadSystemEnv()

	showHeaders := !cfg.NoHeaders
	model := ui.NewModel(parsedFile, envVars, showHeaders)

	p := tea.NewProgram(model, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error running program: %v\n", err)
		os.Exit(1)
	}
}
