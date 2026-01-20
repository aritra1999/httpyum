#!/bin/bash
set -e

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

confirm_uninstall() {
    echo ""
    print_warning "This will remove $BINARY_NAME from your system"
    echo ""
    read -p "Are you sure you want to continue? [y/N] " -n 1 -r
    echo ""

    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        print_info "Uninstall cancelled"
        exit 0
    fi
}

check_installation() {
    BINARY_PATH=$(command -v "$BINARY_NAME" 2>/dev/null || true)

    if [ -z "$BINARY_PATH" ]; then
        print_error "$BINARY_NAME is not installed or not in PATH"
        exit 1
    fi

    print_info "Found $BINARY_NAME at: $BINARY_PATH"

    CURRENT_VERSION=$("$BINARY_NAME" --version 2>&1 || echo "unknown")
    print_info "Version: $CURRENT_VERSION"
}

remove_binary() {
    BINARY_PATH=$(command -v "$BINARY_NAME" 2>/dev/null)

    if [ ! -f "$BINARY_PATH" ]; then
        print_error "Binary not found at $BINARY_PATH"
        exit 1
    fi

    print_info "Removing $BINARY_PATH..."

    BINARY_DIR=$(dirname "$BINARY_PATH")

    if [ -w "$BINARY_DIR" ]; then
        rm -f "$BINARY_PATH"
    else
        print_warning "Removing with sudo (password may be required)..."
        sudo rm -f "$BINARY_PATH"
    fi

    print_success "Removed $BINARY_NAME"
}

verify_removal() {
    print_info "Verifying removal..."

    if command -v "$BINARY_NAME" &> /dev/null; then
        print_warning "$BINARY_NAME is still found in PATH"
        print_info "You may have multiple installations. Check:"
        echo "  which -a $BINARY_NAME"
    else
        print_success "Successfully uninstalled $BINARY_NAME"
    fi
}

cleanup_note() {
    echo ""
    print_info "Note: Dependencies (jless/fx) were not removed"
    echo ""
    echo "If you want to remove them as well:"
    echo ""

    OS=$(uname -s | tr '[:upper:]' '[:lower:]')
    case "$OS" in
        darwin*)
            echo "  brew uninstall jless"
            echo "  brew uninstall fx"
            ;;
        linux*)
            if command -v apt-get &> /dev/null; then
                echo "  sudo apt-get remove jless"
            elif command -v yum &> /dev/null; then
                echo "  sudo yum remove jless"
            elif command -v pacman &> /dev/null; then
                echo "  sudo pacman -R jless"
            fi
            echo "  # For fx, remove the Go binary if installed via 'go install'"
            ;;
        mingw*|msys*|cygwin*)
            echo "  scoop uninstall jless"
            echo "  scoop uninstall fx"
            ;;
    esac
    echo ""
}

main() {
    echo ""
    print_info "httpyum uninstaller"
    echo ""

    check_installation
    confirm_uninstall

    echo ""
    remove_binary
    verify_removal
    cleanup_note

    print_success "Goodbye!"
    echo ""
}

main "$@"
