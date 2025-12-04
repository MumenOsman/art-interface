# Art Interface & RLE Encoder

A versatile tool designed to generate ASCII art and perform custom Run-Length Encoding (RLE) compression and decompression. Whether you prefer a visual web interface or a command-line tool, this project has you covered.

## Features

- **ASCII Art Generation**: Convert standard text into stylish ASCII art (Web Interface).
- **Custom RLE Compression**: Encode strings and ASCII art into a compressed format `[Count Pattern]` to save space.
- **Decompression**: Decode compressed strings back to their original form.
- **Dual Interface**:
    - **Web UI**: A modern, responsive dark-mode interface for easy interaction.
    - **CLI**: A robust command-line tool for quick scripting and processing.

## Getting Started

### Prerequisites

- Go 1.16 or higher

### 1. Web Interface

The web interface provides the complete experience, allowing you to generate art, encode, and decode with a click.

**Start the server:**
```bash
go run . --web
```

**Open in browser:**
Navigate to [http://localhost:8080](http://localhost:8080).

**Web Modes:**
- **Art**: Type text and generate ASCII art.
- **Encode**: Paste text/art to compress it.
- **Decode**: Paste a compressed string to restore it.

### 2. Command Line Interface (CLI)

Use the tool directly from your terminal for encoding and decoding tasks.

**Encode a string:**
Use the `--Encode` flag.
```bash
go run . --Encode "Hello World"
```
*Output: `He[2 l]o World` (Compressed representation)*

**Decode a string:**
Simply pass the encoded string as an argument.
```bash
go run . "He[2 l]o"
```
*Output: `Hello`*

**Multi-line Support:**
Add the `--Multi` flag for processing multi-line strings (useful for large ASCII art blocks).

## Algorithm

The project uses a custom Run-Length Encoding (RLE) algorithm:
- **Format**: `[Count Pattern]` for repetitions > 1. Single characters are left as-is.
- **Example**: `AAAABBB` becomes `[4 A][3 B]`
- **Efficiency**: It intelligently looks for repeating patterns (not just single characters) to maximize compression.

## Design

The web interface features a "Retro Coding" aesthetic with a Gruvbox-inspired dark theme, ensuring a comfortable and stylish user experience.
