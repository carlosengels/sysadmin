# Pwgen - Secure Password Generator

Pwgen is a secure password generator that creates cryptographically secure passwords with a specified length. It ensures that generated passwords contain a mix of lowercase, uppercase, numbers, and special characters.

## Features

- Cryptographically secure random number generation
- Customizable password length
- Ensures password complexity by including:
  - Lowercase letters
  - Uppercase letters
  - Numbers
  - Special characters
- Minimum password length of 8 characters
- No external dependencies

## Prerequisites

- Linux-based operating system:
  - Ubuntu/Debian
  - CentOS/RHEL
  - Fedora
  - Amazon Linux
  - Amazon Linux 2
- Root/sudo privileges
- Internet connection (for initial setup)

The setup script will automatically install:
- Go (if not present)
- Development tools
- Git and curl

## Installation

1. Clone or download this repository:
   ```bash
   git clone https://github.com/carlosengels/sysadmin
   cd sysadmin/pwgen
   ```

2. Make the setup script executable:
   ```bash
   chmod +x setup.sh
   ```

3. Run the setup script with sudo:
   ```bash
   sudo ./setup.sh
   ```

## Usage

### Basic Usage

```bash
pwgen <length>
```

Example:
```bash
pwgen 16
```

### Example Output

```
Kj9#mP2$vL5nX8@q
```

### Password Requirements

- Minimum length: 8 characters
- Contains at least:
  - One lowercase letter
  - One uppercase letter
  - One number
  - One special character
- Uses cryptographically secure random number generation

### Character Sets Used

- Lowercase letters: a-z
- Uppercase letters: A-Z
- Numbers: 0-9
- Special characters: !@#$%^&*()_+-=[]{}|;:,.<>?

## Security Features

1. **Cryptographic Security**
   - Uses `crypto/rand` for cryptographically secure random number generation
   - No predictable patterns in password generation

2. **Complexity Requirements**
   - Guaranteed mix of character types
   - Random distribution of characters
   - No sequential patterns

3. **No External Dependencies**
   - Self-contained implementation
   - No network calls for password generation
   - No external random number sources

## Examples

Generate a 12-character password:
```bash
pwgen 12
```

Generate a 32-character password:
```bash
pwgen 32
```

## Troubleshooting

1. **Permission Denied**
   ```bash
   sudo chmod +x setup.sh
   sudo chmod +x /usr/local/bin/pwgen
   ```

2. **Go Not Found**
   - The setup script should install Go automatically
   - If it fails, install Go manually from https://golang.org/dl/

3. **Build Issues**
   ```bash
   cd pwgen
   go mod tidy
   go build
   ```

## Contributing

Feel free to submit issues, fork the repository, and create pull requests for any improvements.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 