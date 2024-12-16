# seddle

Seddle is a command-line tool that updates text files using natural language instructions powered by OpenAI's GPT-4 model.

## Description

Seddle makes it easy to modify text files using plain English instructions. Instead of writing complex sed commands or regular expressions, you can simply describe what changes you want to make to your file, and seddle will use OpenAI's GPT-4 to interpret and apply those changes.

## Installation

```bash
go install github.com/yourusername/seddle@latest
```

## Prerequisites

- Go 1.23 or higher
- OpenAI API key

## Usage

```bash
# Using command line flags
seddle -f path/to/file.txt -i "your instruction" -k your-api-key

# Using environment variable for API key
export OPENAI_API_KEY=your-api-key
seddle -f path/to/file.txt -i "your instruction"
```

### Examples

```bash
# Add a new section to a markdown file
seddle -f README.md -i "Add a Usage section after the Installation section"

# Update configuration in a JSON file
seddle -f config.json -i "Change the port number to 3000"

# Modify text content
seddle -f document.txt -i "Replace all instances of 'customer' with 'client'"
```

## Features

- Natural language instructions for file modifications
- Supports any text file format
- Creates backup files with `.updated` extension
- Environment variable support for API key
- Simple command-line interface

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with [OpenAI API](https://openai.com/api/)
- Uses [Cobra](https://github.com/spf13/cobra) for CLI interface