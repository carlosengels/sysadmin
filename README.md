# SysAdmin Tools Collection

A collection of powerful system administration and network tools created with Cursor.

## 🛠️ Available Tools

### Network Tools

#### [Geotrace](geotrace/README.md)
A powerful network diagnostic tool that combines traceroute functionality with IP geolocation.
- Traditional traceroute using ICMP packets
- Geolocation information for each hop
- DNS resolution
- ISP information
- Maximum 30 hops (configurable)

### Security Tools

#### [Pwgen](pwgen/README.md)
A secure password generator that creates cryptographically secure passwords.
- Cryptographically secure random number generation
- Customizable password length
- Ensures password complexity
- No external dependencies
- Minimum 8 characters

### S3 Tools

#### [S3Sync](s3sync/README.md)
A simple command-line tool for synchronizing files between your local system and Amazon S3.
- Silent operation by default
- Bi-directional sync (local ↔ S3)
- Uses exact timestamps for accurate sync
- Configurable through environment variables
- Simple one-command operation

## 🚀 Quick Start

1. Clone the repository:
   ```bash
   git clone https://github.com/carlosengels/sysadmin
   cd sysadmin
   ```

2. Navigate to the desired tool directory:
   ```bash
   cd <tool-name>  # e.g., cd geotrace, cd pwgen, or cd s3sync
   ```

3. Follow the tool-specific installation instructions in each tool's README.

## 📋 Requirements

- Linux-based operating system (Ubuntu, Debian, CentOS, RHEL, Fedora, Amazon Linux)
- Root/sudo privileges
- Go (will be installed automatically by setup scripts)
- Internet connection (for initial setup)

## 🔧 Installation

Each tool has its own setup script that will:
- Install required dependencies
- Set up the Go environment
- Build the tool
- Create system-wide executables

## 📚 Documentation

Detailed documentation for each tool can be found in their respective directories:
- [Geotrace Documentation](geotrace/README.md)
- [Pwgen Documentation](pwgen/README.md)
- [S3Sync Documentation](s3sync/README.md)

## 🤝 Contributing

Feel free to submit issues, fork the repository, and create pull requests for any improvements.

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

Made with ❤️ using [Cursor](https://cursor.sh)