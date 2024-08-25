# ls-clone

`ls-clone` is a command-line tool written in Go that mimics the Unix `ls` command. It provides options to display files in a long format, reverse order, and include hidden files (dotfiles).

## Features

- **Long Format Display**: Shows detailed information about each file.
- **Reverse Order**: Displays files in reverse sorted order.
- **All Files**: Includes hidden files in the output.

## Installation

To install:

```bash
go install github.com/yourusername/ls-clone@latest
```

Or, clone the repo and build manually:

```bash
git clone https://github.com/yourusername/ls-clone.git
cd ls-clone
go build -o ls-clone
```

## Usage

Run `ls-clone` with the following options:

- `-l`: Long format
- `-r`: Reverse order
- `-a`: Include hidden files

## License

This project is open-source and available under the MIT License.
