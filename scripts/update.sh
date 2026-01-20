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

get_current_version() {
    if ! command -v "$BINARY_NAME" &> /dev/null; then
        print_error "$BINARY_NAME is not installed"
        print_info "Run the install script instead:"
        echo "  curl -fsSL https://raw.githubusercontent.com/$REPO/main/scripts/install.sh | bash"
        exit 1
    fi

    CURRENT_VERSION=$("$BINARY_NAME" --version 2>&1 | grep -oE 'v[0-9]+\.[0-9]+\.[0-9]+' || echo "unknown")
    print_info "Current version: $CURRENT_VERSION"
}

get_latest_version() {
    print_info "Fetching latest version..."
    LATEST_VERSION=$(curl -fsSL "https://api.github.com/repos/$REPO/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')

    if [ -z "$LATEST_VERSION" ]; then
        print_error "Failed to fetch latest version"
        exit 1
    fi

    print_info "Latest version: $LATEST_VERSION"
}

check_if_update_needed() {
    if [ "$CURRENT_VERSION" = "$LATEST_VERSION" ]; then
        print_success "Already up to date!"
        exit 0
    fi

    print_info "Update available: $CURRENT_VERSION → $LATEST_VERSION"
}

download_and_install() {
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

    INSTALL_PATH="$INSTALL_DIR/$BINARY_NAME"

    print_info "Updating $INSTALL_PATH..."

    if [ -w "$INSTALL_DIR" ]; then
        mv "$BINARY_FILE" "$INSTALL_PATH"
    else
        print_warning "Updating with sudo (password may be required)..."
        sudo mv "$BINARY_FILE" "$INSTALL_PATH"
    fi

    print_success "Updated $BINARY_NAME to $INSTALL_PATH"
}

verify_update() {
    print_info "Verifying update..."

    UPDATED_VERSION=$("$BINARY_NAME" --version 2>&1 | grep -oE 'v[0-9]+\.[0-9]+\.[0-9]+' || echo "unknown")

    if [ "$UPDATED_VERSION" = "$LATEST_VERSION" ]; then
        print_success "Update successful!"
        echo ""
        echo "Updated: $CURRENT_VERSION → $UPDATED_VERSION"
    else
        print_warning "Update completed but version verification failed"
        echo "Expected: $LATEST_VERSION"
        echo "Got: $UPDATED_VERSION"
    fi
}

main() {
    echo ""
    print_info "httpyum updater"
    echo ""

    detect_os
    detect_arch
    get_current_version
    get_latest_version
    check_if_update_needed

    echo ""
    download_and_install "$LATEST_VERSION" "$OS" "$ARCH"
    verify_update

    echo ""
    print_success "All done!"
    echo ""
}

main "$@"
