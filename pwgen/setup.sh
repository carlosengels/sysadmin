#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to print status messages
print_status() {
    echo -e "${GREEN}[*]${NC} $1"
}

print_error() {
    echo -e "${RED}[!]${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}[!]${NC} $1"
}

# Check if running as root
if [ "$EUID" -ne 0 ]; then
    print_error "This script must be run as root (use sudo)"
    exit 1
fi

# Enhanced OS detection
detect_os() {
    if [ -f /etc/os-release ]; then
        . /etc/os-release
        OS=$NAME
        VERSION_ID=$VERSION_ID
        
        # Special handling for Amazon Linux
        if [[ $OS == "Amazon Linux" ]] || grep -q "Amazon Linux" /etc/system-release 2>/dev/null; then
            OS="Amazon Linux"
            print_status "Detected Amazon Linux"
            return
        fi
    elif [ -f /etc/system-release ]; then
        OS=$(cat /etc/system-release | cut -d' ' -f1)
    elif [ -f /etc/lsb-release ]; then
        . /etc/lsb-release
        OS=$DISTRIB_ID
    else
        OS=$(uname -s)
    fi
    print_status "Detected OS: $OS"
}

# Install Go based on OS
install_go() {
    print_status "Installing Go..."
    
    case $OS in
        "Ubuntu"|"Debian GNU/Linux")
            apt-get update
            apt-get install -y golang
            ;;
        "Amazon Linux")
            # Amazon Linux specific installation
            amazon-linux-extras install -y golang1.11
            yum install -y golang
            ;;
        "CentOS Linux"|"Red Hat Enterprise Linux")
            yum install -y golang
            ;;
        "Fedora")
            dnf install -y golang
            ;;
        *)
            print_error "Unsupported operating system for automatic Go installation"
            print_warning "Please install Go manually from https://golang.org/dl/"
            exit 1
            ;;
    esac
}

# Install required system packages
install_prerequisites() {
    print_status "Installing prerequisites..."
    
    case $OS in
        "Ubuntu"|"Debian GNU/Linux")
            apt-get update
            apt-get install -y git curl build-essential
            ;;
        "Amazon Linux")
            yum update -y
            yum groupinstall -y "Development Tools"
            yum install -y git curl
            ;;
        "CentOS Linux"|"Red Hat Enterprise Linux")
            yum update -y
            yum groupinstall -y "Development Tools"
            yum install -y git curl
            ;;
        "Fedora")
            dnf update -y
            dnf groupinstall -y "Development Tools"
            dnf install -y git curl
            ;;
        *)
            print_error "Unsupported operating system for automatic prerequisite installation"
            exit 1
            ;;
    esac
}

# Detect OS
detect_os

# Install prerequisites
install_prerequisites

# Check if Go is installed and install if needed
if ! command -v go &> /dev/null; then
    print_warning "Go is not installed"
    install_go
fi

# Verify Go installation
if ! command -v go &> /dev/null; then
    print_error "Go installation failed. Please install Go manually."
    exit 1
fi

# Print Go version
go version

# Create project directory if it doesn't exist
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
print_status "Setting up project in $SCRIPT_DIR"

# Initialize Go module
if [ ! -f "$SCRIPT_DIR/go.mod" ]; then
    print_status "Initializing Go module..."
    cd "$SCRIPT_DIR"
    go mod init pwgen
fi

# Build the program
print_status "Building pwgen..."
go build -o pwgen

# Create a symbolic link to the binary
ln -sf "$SCRIPT_DIR/pwgen" /usr/local/bin/pwgen

print_status "Installation complete!"
print_status "You can now run the program using: pwgen <length>"
print_status "Example: pwgen 16"

# Check if the installation was successful
if [ -f "$SCRIPT_DIR/pwgen" ]; then
    print_status "Setup completed successfully!"
else
    print_error "Setup failed. Please check the error messages above."
    exit 1
fi 