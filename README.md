# httpyum


![Loading Preview ...](https://github.com/user-attachments/assets/fc453cdb-5173-48a9-9455-3551bb1ab805)


A fast, interactive CLI tool for executing HTTP requests from `.http` files, built with Go and [bubbletea](https://github.com/charmbracelet/bubbletea).

## Features

- Parse and execute HTTP requests from `.http` files
- **Beautiful interactive TUI with fuzzy search filtering** for selecting and running requests
- Support for variables and variable substitution
- Elegant response display with syntax highlighting
- **Interactive JSON exploration with [jless](https://jless.io/)** - press `f` to expand/collapse and explore JSON responses interactively
- Support for multiple HTTP methods (GET, POST, PUT, DELETE, PATCH, etc.)
- Request headers and body support
- Response timing and size information
- Toggleable header display
- Fast and lightweight

## Installation

### Quick Install (Recommended)

Install httpyum and its dependencies with a single command:

```bash
curl -fsSL https://raw.githubusercontent.com/aritra1999/httpyum/main/scripts/install.sh | bash
```

This will:
- Download the latest release for your platform
- Install httpyum to `/usr/local/bin`
- Provide instructions for installing jless (recommended for JSON viewing)

### Manual Installation

Download the latest release for your platform from the [releases page](https://github.com/aritra1999/httpyum/releases):

**macOS (Apple Silicon)**
```bash
curl -L https://github.com/aritra1999/httpyum/releases/latest/download/httpyum-darwin-arm64.tar.gz | tar xz
sudo mv httpyum-darwin-arm64 /usr/local/bin/httpyum
```

**macOS (Intel)**
```bash
curl -L https://github.com/aritra1999/httpyum/releases/latest/download/httpyum-darwin-amd64.tar.gz | tar xz
sudo mv httpyum-darwin-amd64 /usr/local/bin/httpyum
```

**Linux (amd64)**
```bash
curl -L https://github.com/aritra1999/httpyum/releases/latest/download/httpyum-linux-amd64.tar.gz | tar xz
sudo mv httpyum-linux-amd64 /usr/local/bin/httpyum
```

**Windows**
Download from [releases page](https://github.com/aritra1999/httpyum/releases) and add to PATH.

### Install from Source

```bash
go build -o httpyum ./cmd/httpyum
```

Or install directly:

```bash
go install httpyum/cmd/httpyum@latest
```

## Updating

Update to the latest version:

```bash
curl -fsSL https://raw.githubusercontent.com/aritra1999/httpyum/main/scripts/update.sh | bash
```

## Uninstalling

Remove httpyum from your system:

```bash
curl -fsSL https://raw.githubusercontent.com/aritra1999/httpyum/main/scripts/uninstall.sh | bash
```

## Dependencies

### Optional: Install jless for Interactive JSON Viewing

To enable interactive JSON exploration with expand/collapse, install [jless](https://jless.io/) (recommended):

```bash
# macOS
brew install jless

# Linux (cargo required)
cargo install jless

# Or download from: https://github.com/PaulJuliusMartinez/jless/releases
```

Alternatively, install [fx](https://fx.wtf/):

```bash
# macOS
brew install fx

# Linux / npm / Go
curl -sS https://webi.sh/fx | sh
npm install -g fx
go install github.com/antonmedv/fx@latest
```

When viewing a JSON response, press `f` to open it interactively. httpyum will use jless if available (better expand/collapse), otherwise fx.

**jless keybindings (recommended):**
- `↑`/`↓` or `j`/`k` - Navigate through keys
- `Enter` or `Space` - **Expand/collapse** objects and arrays
- `e` - Expand all
- `c` - Collapse all
- `/` - Search
- `n`/`N` - Next/previous search result
- `%` - Toggle quotes on string values
- `q` - Quit and return to httpyum

## Usage

```bash
httpyum [OPTIONS] <file.http>
```

### Options

- `--no-headers` - Hide response headers in output
- `-h, --help` - Show help message
- `-v, --version` - Show version information

### Examples

```bash
# Run with example file
httpyum example.http

# Run without showing headers
httpyum --no-headers api.http
```

## Keyboard Controls

### List View
- `↑`/`↓` or `k`/`j` - Navigate requests
- `/` - Filter requests (fuzzy search)
- `Esc` - Clear filter
- `Enter` - Execute selected request
- `q` - Quit

### Response View
- `f` - Open JSON response in interactive viewer (jless/fx) with expand/collapse (JSON responses only)
- `h` - Toggle headers visibility
- `v` - Toggle variables panel (shows variables used in request)
- `b` or `Esc` - Back to list
- `q` - Quit

## .http File Format

httpyum supports the standard `.http` file format:

```http
# Comments start with # or //
@variable = value

### Request Description
METHOD https://api.example.com/{{variable}}
Header-Name: Header-Value
Another-Header: Value

{
  "body": "json"
}

### Another Request
GET https://example.com
```

### Variables

Define variables with `@name = value` and use them with `{{name}}`:

```http
@baseUrl = https://api.example.com
@token = your-api-token

### Get Users
GET {{baseUrl}}/users
Authorization: Bearer {{token}}
```

### Environment Variables

Load environment variables from your shell using `{{$dotenv VARIABLE_NAME}}`:

**Set environment variables in your shell:**
```bash
export JWT="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
export API_KEY="your-api-key-here"
```

**Use in your `.http` file:**
```http
@baseUrl = http://localhost:3000
@jwt = {{$dotenv JWT}}

### Authenticated Request
GET {{baseUrl}}/api/users
Authorization: Bearer {{jwt}}

### Or use directly
POST {{baseUrl}}/api/data
X-API-Key: {{$dotenv API_KEY}}
Content-Type: application/json

{
  "data": "example"
}
```

**Features:**
- ✅ Reads environment variables from your shell
- ✅ Shows only variables used in the request (press `v` to toggle)
- ✅ Masked values for security (only shows last 3 characters)
- ✅ Variables can reference env vars: `@token = {{$dotenv JWT}}`

### Request Separators

Requests are separated by `###` optionally followed by a description:

```http
### Get user profile
GET https://api.example.com/profile

### Update user profile
POST https://api.example.com/profile
Content-Type: application/json

{
  "name": "John Doe"
}
```

## Example .http File

See [example.http](./example.http) for a complete example with various request types.

## Features Supported

- ✅ HTTP Methods (GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS, etc.)
- ✅ Request headers
- ✅ Request body (JSON, form data, text)
- ✅ Variables and variable substitution
- ✅ Environment variables (shell environment with `{{$dotenv VAR}}`)
- ✅ Comments (`#` and `//`)
- ✅ Request descriptions
- ✅ Response display with timing
- ✅ JSON pretty-printing
- ✅ Toggleable headers
- ✅ Status code colorization

## Why httpyum?

httpyum was created as a faster alternative to [httpyac](https://httpyac.github.io/). Built with Go and featuring an interactive TUI, httpyum provides:

- Instant startup time
- Low memory footprint
- Clean, intuitive interface
- Fast request execution
- Cross-platform support

## Building from Source

```bash
# Clone the repository
git clone https://github.com/aritra1999/httpyum.git
cd httpyum

# Install dependencies
go mod download

# Build
go build -o httpyum ./cmd/httpyum

# Run
./httpyum example.http
```

## Project Structure

```
httpyum/
├── cmd/httpyum/          # Main CLI entry point
├── internal/
│   ├── parser/           # .http file parsing
│   ├── client/           # HTTP request execution
│   ├── ui/               # Bubbletea TUI
│   └── config/           # CLI configuration
├── example.http          # Example requests
└── README.md
```

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.
