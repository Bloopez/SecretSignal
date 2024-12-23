# DNS over HTTPS Backdoor

A secure command execution tool that leverages DNS TXT records and HTTPS for covert communication.

## Overview
This tool enables secure command execution by retrieving encrypted payloads from DNS TXT records over HTTPS. The two-layer encryption (AES + TLS) ensures secure transmission of commands while maintaining a low network footprint.

## Key Features
- **Double Encryption Protection**
  - AES encryption for payload content
  - TLS encryption for data transmission
- **DNS Covert Channel**
  - Utilizes DNS TXT records for command delivery
  - Minimizes suspicious network traffic
- **Windows Compatible**
  - Optimized for Windows environments
  - Clean PowerShell execution

## Components

### generator.go
- Encrypts command payloads using AES
- Prepares data for DNS TXT record storage
- Manages payload generation workflow

### main.go
- Establishes secure HTTPS connections
- Retrieves encrypted DNS TXT records
- Decrypts and executes PowerShell commands

## Usage Instructions

1. Configure your domain in `main.go`
2. Set your command in `generator.go` under `CmdContent`
3. Compile for Windows:
```bash
GOOS=windows GOARCH=amd64 go build main.go
```

## Requirements
- Go development environment
- Windows target system
- Domain with configurable DNS TXT records

## Important Notes
- Currently supports Windows systems only
- Ensure proper DNS record configuration in `main.go`

## Contributing
Contributions welcome! Please follow standard Go coding conventions and include tests for new features.

---
**Disclaimer**: Use responsibly and in compliance with applicable laws and regulations.
