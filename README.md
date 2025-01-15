# opml2csv

`opml2csv` is a command-line tool written in Go for converting OPML (Outline Processor Markup Language) files into CSV files. This tool extracts feed information defined in an OPML file and structures it into a tabular, CSV-friendly format.

## Features

- Efficient conversion of an OPML file to CSV.
- Command-line arguments that make input and output handling simple.
- Automatic generation of CSV headers.
- Validation of OPML input files.
- Options to display help and version information.

## Prerequisites

- **Go SDK 1.23.4** (or a more recent version).
- An OPML file containing structured feed information (feeds, URLs, etc.).

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/Kharec/opml2csv.git
   cd opml2csv
   ```

2. Build the project:

   ```bash
   go build -o opml2csv
   ```

3. Push it into a bin directory in your `$PATH`:

   ```bash
   mv ompl2csv ~/bin # for example
   ```

## Usage

### Command Syntax

```bash
opml2csv [OPTIONS]
```

### Options

| Option           | Description                                                      |
|-------------------|------------------------------------------------------------------|
| `-h`, `--help`    | Displays help with usage information.                           |
| `-v`, `--version` | Displays the current version of the tool.                       |
| `-i`, `--input`   | Path to the OPML file to be converted.                          |
| `-o`, `--output`  | Path to the generated CSV file (default: `./export.csv`).       |

### Examples

1. Display help:

   ```bash
   ./opml2csv -h
   ```

2. Convert an OPML file (`feeds.opml`) to CSV (`feeds.csv`):

   ```bash
   ./opml2csv -i feeds.opml -o feeds.csv
   ```

3. Display the version of the tool:

   ```bash
   ./opml2csv -v
   ```

## Project Structure

The project is organized as follows:

```plaintext
opml2csv/
├── internal/
│   ├── flags/            # Handles command-line arguments
│   │   ├── parse.go      # Parsing of options (input, output, etc.)
│   │   ├── manage.go     # Functions for handling help, version, validation
│   │   ├── help.go       # Function to display the help text
│   └── utils/            # Utilities for file processing
│       ├── parsers.go    # Parses OPML lines into Outline structs
│       ├── builders.go   # Builds CSV headers and lines
│       ├── types.go      # Defines structured types (e.g., Outline)
│       └── version.go    # Manages the application version
├── opml2csv.go           # Main code (entry point for the application)
```

## How It Works

1. **Reading the OPML File**: OPML files contain structured lists of feeds in XML format.
2. **Parsing**: Relevant OPML lines containing structured data are extracted.
3. **CSV Validation and Generation**: Feed information is processed and converted into a cleanly formatted CSV file.

## Contributing

Contributions, bug reports, and feature requests are welcome! Feel free to submit an **issue** or a **pull request** on the GitHub repository.

## License

This project is licensed under the [MIT License](LICENSE). You are free to use, modify, and distribute the project.
