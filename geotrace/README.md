# Geotrace - Enhanced Traceroute with Geolocation

Geotrace is a powerful network diagnostic tool that combines the functionality of traceroute with IP geolocation. It traces the route of packets from your computer to any destination on the Internet while providing detailed geographic information about each hop along the way.

## Features

- Traditional traceroute functionality using ICMP packets
- Geolocation information for each hop including:
  - City and Country
  - ISP (Internet Service Provider)
  - Geographic coordinates
- DNS resolution for hostnames
- Colorized and formatted output
- Maximum 30 hops (configurable)
- Timeout of 2 seconds per probe

## Prerequisites

- Linux-based operating system:
  - Ubuntu/Debian
  - CentOS/RHEL
  - Fedora
  - Amazon Linux
  - Amazon Linux 2
- Root/sudo privileges
- Internet connection

The setup script will automatically install:
- Go (if not present)
- Development tools
- Git and curl
- Required Go dependencies

## Installation

1. Clone or download this repository:
   ```bash
   git clone https://github.com/carlosengels/sysadmin
   cd sysadmin/geotrace
   ```

2. Make the setup script executable:
   ```bash
   chmod +x setup.sh
   ```

3. Run the setup script with sudo:
   ```bash
   sudo ./setup.sh
   ```

The setup script will:
- Check and install Go if not present
- Install required dependencies
- Build the geotrace program
- Create a system-wide executable
- Set up proper permissions

## Usage

### Basic Usage

```bash
sudo geotrace <hostname or IP>
```

Example:
```bash
sudo geotrace google.com
```

### Example Output

```
Tracing route to google.com [172.217.xxx.xxx]
Maximum hops: 30

 1  192.168.1.1  1.234 ms  [New York, United States (US) - Comcast Cable]
 2  10.0.0.1 (router.isp.net)  15.678 ms  [Chicago, United States (US) - Level 3 Communications]
 3  172.16.0.1  25.901 ms  [Dallas, United States (US) - Google LLC]
...
```

### Understanding the Output

Each line in the output contains:
1. Hop number
2. IP address and/or hostname (if available)
3. Response time in milliseconds
4. Geographic location (City, Country)
5. Internet Service Provider (ISP)

### Error Messages

- `* * *`: Indicates no response was received from this hop
- `Could not resolve <hostname>`: DNS resolution failed for the given hostname
- `This program must be run as root`: You need to use sudo to run the program

## Troubleshooting

1. **Permission Denied**
   ```bash
   sudo chmod +x setup.sh
   sudo chmod +x /usr/local/bin/geotrace
   ```

2. **Go Not Found**
   - The setup script should install Go automatically
   - If it fails, install Go manually from https://golang.org/dl/

3. **Dependencies Issues**
   ```bash
   cd geotrace
   go mod tidy
   go mod download
   ```

4. **Rate Limiting**
   - The geolocation service (ip-api.com) has a rate limit of 45 requests per minute
   - The program includes a built-in delay to respect these limits

## Technical Details

- Uses ICMP Echo Request/Reply packets
- Default packet size: 52 bytes
- Maximum number of hops: 30 (configurable in source)
- Timeout per probe: 2 seconds
- Geolocation data provided by ip-api.com
- Written in Go for high performance and reliability

## Limitations

1. Requires root/sudo privileges due to raw socket usage
2. Some routers may block ICMP packets
3. Private IP addresses won't return geolocation data
4. Rate limiting on the geolocation API (45 requests/minute)

## Security Considerations

- The program requires root privileges due to its use of raw sockets
- No sensitive data is stored or transmitted
- Uses HTTP for geolocation queries (ip-api.com free tier limitation)

## Contributing

Feel free to submit issues, fork the repository, and create pull requests for any improvements.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 