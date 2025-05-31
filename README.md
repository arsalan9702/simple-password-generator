# Password Generator

A simple and customizable command-line password generator written in Go using the Cobra CLI framework.

## Features

- Generate passwords with customizable length
- Include/exclude digits (0-9)
- Include/exclude special characters
- Simple command-line interface
- Built with Go and Cobra CLI framework

## Installation

### Prerequisites

- Go 1.23.4 or higher

### Install from source

1. Clone the repository:
```bash
git clone <repository-url>
cd passwordGenerator
```

2. Build the application:
```bash
go build -o passwordGenerator
```

3. (Optional) Install globally:
```bash
go install
```

## Usage

### Basic Usage

Generate a default 8-character password (letters only):
```bash
./passwordGenerator generate
```

### Advanced Usage

Generate a password with custom options:

```bash
# Generate a 12-character password with digits
./passwordGenerator generate -l 12 -d

# Generate a 16-character password with digits and special characters
./passwordGenerator generate -l 16 -d -s

# Generate a 20-character password with special characters only
./passwordGenerator generate --length 20 --special-char
```

### Command Options

| Flag | Short | Description | Default |
|------|-------|-------------|---------|
| `--length` | `-l` | Length of the password | 8 |
| `--digits` | `-d` | Include digits (0-9) | false |
| `--special-char` | `-s` | Include special characters | false |

### Help

View available commands and options:
```bash
./passwordGenerator --help
./passwordGenerator generate --help
```

## Character Sets

- **Letters**: A-Z, a-z (always included)
- **Digits**: 0-9 (included with `-d` flag)
- **Special Characters**: `!@#$%^&*()_+{}[]|;:<>,.?-=` (included with `-s` flag)

## Examples

```bash
# Simple 8-character password
./passwordGenerator generate
# Output: AbCdEfGh

# 12-character password with numbers
./passwordGenerator generate -l 12 -d
# Output: AbC3EfG8iJkL

# 16-character password with numbers and special characters
./passwordGenerator generate -l 16 -d -s
# Output: A3c$E7g!i2kL@nO5
```

## Development

### Project Structure

```
passwordGenerator/
├── cmd/
│   ├── password.go    # Password generation command
│   └── root.go        # Root command configuration
├── main.go            # Application entry point
├── go.mod             # Go module file
├── go.sum             # Go dependencies
└── README.md          # This file
```

### Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework for Go

### Building

```bash
# Build for current platform
go build

# Build for different platforms
GOOS=linux GOARCH=amd64 go build -o passwordGenerator-linux
GOOS=windows GOARCH=amd64 go build -o passwordGenerator.exe
GOOS=darwin GOARCH=amd64 go build -o passwordGenerator-mac
```

### Running Tests

```bash
go test ./...
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Security Note

This password generator uses Go's `math/rand` package, which is suitable for general use but not cryptographically secure. For cryptographically secure passwords, consider using `crypto/rand` instead.

## License

This project is open source. Please add your preferred license.

## Author

Add your name and contact information here.
