# Art Decoder Interface

Welcome to the Art Decoder Interface! This tool allows you to encode and decode ASCII art using a beautiful web interface or the command line.

## Features

- **Web Interface**: A modern, responsive web UI to interact with the encoder/decoder.
- **CLI Support**: Retains the original command-line functionality.
- **Multi-line Support**: Handle complex ASCII art with ease.
- **Real-time Feedback**: See the results instantly.

## Getting Started

### Prerequisites

- Go 1.16 or higher

### Running the Web Interface

To start the web server, run the following command from the project root:

```bash
go run . --web
```

Then, open your browser and navigate to:

[http://localhost:8080](http://localhost:8080)

### Using the CLI

You can still use the command line interface as before:

**Encode:**
```bash
go run . --Encode "Hello"
```

**Decode:**
```bash
go run . "EncodedString"
```

## Usage

1.  **Enter Text**: Type or paste your text into the input area.
2.  **Select Mode**: Choose "Encode" to turn text into art, or "Decode" to revert it.
3.  **Generate**: Click the button to see the magic happen!

## Design

The interface is designed with a focus on aesthetics, featuring a dark mode theme, smooth animations, and a clean layout to inspire creativity.

Enjoy creating!
