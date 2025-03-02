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

# Detect OS
if [ -f /etc/os-release ]; then
    . /etc/os-release
    OS=$NAME
elif [ -f /etc/lsb-release ]; then
    . /etc/lsb-release
    OS=$DISTRIB_ID
else
    OS=$(uname -s)
fi

# Install Go based on OS
install_go() {
    print_status "Installing Go..."
    
    case $OS in
        "Ubuntu"|"Debian GNU/Linux")
            apt-get update
            apt-get install -y golang
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

# Check if Go is installed
if ! command -v go &> /dev/null; then
    print_warning "Go is not installed"
    install_go
fi

# Create project directory if it doesn't exist
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
print_status "Setting up project in $SCRIPT_DIR"

# Initialize Go module
if [ ! -f "$SCRIPT_DIR/go.mod" ]; then
    print_status "Initializing Go module..."
    cd "$SCRIPT_DIR"
    go mod init geotrace
fi

# Install required dependencies
print_status "Installing dependencies..."
go get golang.org/x/net/icmp
go get golang.org/x/net/ipv4

# Build the program
print_status "Building geotrace..."
go build -o geotrace

# Create a wrapper script for easier execution
cat > /usr/local/bin/geotrace << 'EOF'
#!/bin/bash
if [ "$EUID" -ne 0 ]; then
    echo "This program must be run as root (use sudo)"
    exit 1
fi
SCRIPT_DIR="$(dirname $(readlink -f $(which geotrace)))"
"$SCRIPT_DIR/geotrace" "$@"
EOF

# Make the wrapper script executable
chmod +x /usr/local/bin/geotrace

# Create symbolic link to the binary
ln -sf "$SCRIPT_DIR/geotrace" /usr/local/bin/geotrace

print_status "Installation complete!"
print_status "You can now run the program using: sudo geotrace <hostname>"
print_status "Example: sudo geotrace google.com"

# Check if the installation was successful
if [ -f "$SCRIPT_DIR/geotrace" ]; then
    print_status "Setup completed successfully!"
else
    print_error "Setup failed. Please check the error messages above."
    exit 1
fi 