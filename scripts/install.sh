#!/bin/bash
set -e

REPO="aritra1999/httpyum"
INSTALL_DIR="${INSTALL_DIR:-/usr/local/bin}"
BINARY_NAME="httpyum"

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m'

print_info() {
    echo -e "${BLUE}==>${NC} $1"
}

print_success() {
    echo -e "${GREEN}✓${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}!${NC} $1"
}

print_error() {
    echo -e "${RED}✗${NC} $1"
}

detect_os() {
    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    case "$OS" in
        linux*)
            OS="linux"
            ;;
        darwin*)
            OS="darwin"
            ;;
        mingw*|msys*|cygwin*)
            OS="windows"
            ;;
        *)
            print_error "Unsupported operating system: $OS"
            exit 1
            ;;
    esac
}

detect_arch() {
    ARCH=$(uname -m)
    case "$ARCH" in
        x86_64|amd64)
            ARCH="amd64"
            ;;
        aarch64|arm64)
            ARCH="arm64"
            ;;
        *)
            print_error "Unsupported architecture: $ARCH"
            exit 1
            ;;
    esac
}

get_latest_version() {
    print_info "Fetching latest version..."
    LATEST_VERSION=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

    if [ -z "$LATEST_VERSION" ]; then
        print_error "Failed to fetch latest version"
        exit 1
    fi

    print_success "Latest version: $LATEST_VERSION"
}

download_binary() {
    local version=$1
    local os=$2
    local arch=$3

    if [ "$os" = "windows" ]; then
        ARCHIVE_NAME="${BINARY_NAME}-${os}-${arch}.exe.zip"
        BINARY_FILE="${BINARY_NAME}-${os}-${arch}.exe"
    else
        ARCHIVE_NAME="${BINARY_NAME}-${os}-${arch}.tar.gz"
        BINARY_FILE="${BINARY_NAME}-${os}-${arch}"
    fi

    DOWNLOAD_URL="https://github.com/$REPO/releases/download/$version/$ARCHIVE_NAME"

    print_info "Downloading $ARCHIVE_NAME..."

    TMP_DIR=$(mktemp -d)
    trap "rm -rf $TMP_DIR" EXIT

    if ! curl -fsSL "$DOWNLOAD_URL" -o "$TMP_DIR/$ARCHIVE_NAME"; then
        print_error "Failed to download binary from $DOWNLOAD_URL"
        exit 1
    fi

    print_success "Downloaded successfully"

    print_info "Extracting archive..."
    cd "$TMP_DIR"

    if [ "$os" = "windows" ]; then
        unzip -q "$ARCHIVE_NAME"
    else
        tar -xzf "$ARCHIVE_NAME"
    fi

    print_success "Extracted successfully"

    if [ ! -f "$BINARY_FILE" ]; then
        print_error "Binary file not found in archive"
        exit 1
    fi

    chmod +x "$BINARY_FILE"

    print_info "Installing to $INSTALL_DIR..."

    if [ -w "$INSTALL_DIR" ]; then
        mv "$BINARY_FILE" "$INSTALL_DIR/$BINARY_NAME"
    else
        print_warning "Installing with sudo (password may be required)..."
        sudo mv "$BINARY_FILE" "$INSTALL_DIR/$BINARY_NAME"
    fi

    print_success "Installed $BINARY_NAME to $INSTALL_DIR/$BINARY_NAME"
}

install_dependencies() {
    print_info "Checking dependencies..."

    if command -v jless &> /dev/null; then
        print_success "jless is already installed"
        return
    fi

    print_warning "jless is not installed (recommended for interactive JSON viewing)"
    echo ""
    echo "To install jless:"

    case "$OS" in
        darwin)
            echo "  brew install jless"
            ;;
        linux)
            if command -v apt-get &> /dev/null; then
                echo "  sudo apt-get install jless"
            elif command -v yum &> /dev/null; then
                echo "  sudo yum install jless"
            elif command -v pacman &> /dev/null; then
                echo "  sudo pacman -S jless"
            else
                echo "  cargo install jless  # or download from https://github.com/PaulJuliusMartinez/jless/releases"
            fi
            ;;
        windows)
            echo "  scoop install jless"
            echo "  # or download from https://github.com/PaulJuliusMartinez/jless/releases"
            ;;
    esac

    echo ""
    echo "Alternatively, you can use fx:"
    case "$OS" in
        darwin)
            echo "  brew install fx"
            ;;
        linux)
            echo "  go install github.com/antonmedv/fx@latest"
            ;;
        windows)
            echo "  scoop install fx"
            ;;
    esac
    echo ""
}

verify_installation() {
    print_info "Verifying installation..."

    if ! command -v "$BINARY_NAME" &> /dev/null; then
        print_error "Installation failed: $BINARY_NAME not found in PATH"
        print_warning "You may need to add $INSTALL_DIR to your PATH"
        exit 1
    fi

    VERSION_OUTPUT=$("$BINARY_NAME" --version 2>&1 || true)
    print_success "Installation successful!"
    echo ""
    echo "$VERSION_OUTPUT"
}

main() {
    echo ""
    print_info "httpyum installer"
    echo ""

    detect_os
    detect_arch
    get_latest_version

    print_info "Installing httpyum for $OS/$ARCH"
    echo ""

    download_binary "$LATEST_VERSION" "$OS" "$ARCH"
    install_dependencies
    verify_installation

    echo ""
    print_success "All done! Run 'httpyum --help' to get started"
    echo ""
}

main "$@"
