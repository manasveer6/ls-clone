# ls-clone

`ls-clone` is a simple command-line tool written in Go that mimics the functionality of the Unix `ls` command. It provides options to display files in a long format, reverse order, and include hidden files (dotfiles).

## Features

- **Long Format Display**: Shows detailed information about each file, including name, mode, and human-readable size.
- **Reverse Order**: Displays files in reverse sorted order.
- **All Files**: Option to display all files, including hidden files (dotfiles).

## Installation

To install `ls-clone`, you need to have [Go](https://golang.org/dl/) installed on your system.

```sh
go install github.com/yourusername/ls-clone@latest

```

Or, you can clone the repository and build the binary manually:

```sh
git clone https://github.com/yourusername/ls-clone.git
cd ls-clone
go build -o ls-clone

```
