package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	FilePath    string
	NoHeaders   bool
	ShowHelp    bool
	ShowVersion bool
}

var version = "dev"

func Parse() (*Config, error) {
	cfg := &Config{}

	flag.BoolVar(&cfg.NoHeaders, "no-headers", false, "Hide response headers")
	flag.BoolVar(&cfg.ShowHelp, "help", false, "Show help message")
	flag.BoolVar(&cfg.ShowHelp, "h", false, "Show help message (shorthand)")
	flag.BoolVar(&cfg.ShowVersion, "version", false, "Show version")
	flag.BoolVar(&cfg.ShowVersion, "v", false, "Show version (shorthand)")

	flag.Usage = func() {
		printUsage()
	}

	flag.Parse()
	if cfg.ShowHelp {
		printUsage()
		os.Exit(0)
	}

	if cfg.ShowVersion {
		fmt.Printf("httpyum version %s\n", version)
		os.Exit(0)
	}

	args := flag.Args()
	if len(args) < 1 {
		return nil, fmt.Errorf("missing required argument: file path\n\nUsage: httpyum [OPTIONS] <file.http>")
	}

	cfg.FilePath = args[0]
	if _, err := os.Stat(cfg.FilePath); os.IsNotExist(err) {
		return nil, fmt.Errorf("file not found: %s", cfg.FilePath)
	}

	return cfg, nil
}

func printUsage() {
	fmt.Fprintf(os.Stderr, `httpyum - Fast HTTP request runner for .http files

Usage:
  httpyum [OPTIONS] <file.http>

Arguments:
  <file.http>    Path to .http file containing HTTP requests

Options:
  --no-headers   Hide response headers in output
  -h, --help     Show this help message
  -v, --version  Show version information

Examples:
  httpyum requests.http
  httpyum --no-headers api.http

Keyboard Controls:
  List View:
    ↑/↓          Navigate requests
    /            Filter requests
    Enter        Execute selected request
    q            Quit

  Response View:
    f            Open JSON in interactive viewer (jless/fx)
    h            Toggle headers visibility
    esc/b        Back to list
    q            Quit

.http File Format:
  # Comment
  @variable = value

  ### Request Description
  GET https://api.example.com/{{variable}}
  Header-Name: Header-Value

  {
    "body": "json"
  }

For more information, visit: https://github.com/aritra1999/httpyum
`)
}
