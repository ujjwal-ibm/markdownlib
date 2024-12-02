# MarkdownLib

**MarkdownLib** is a Go library for reading, writing, parsing, and validating Markdown files. It provides an easy-to-use API for developers to programmatically manage Markdown content.

## Features

- Read and write Markdown files.
- Parse Markdown elements such as headings, lists, and tables.
- Validate Markdown structure.
- Append or generate Markdown content.

## Installation

```bash
go get github.com/ujjwal-ibm/markdownlib
```

## Usage

```go
package main

import (
	"log"

	"github.com/ujjwal-ibm/markdownlib"
)

func main() {
	md := markdownlib.NewMarkdownFile("# Title\nContent goes here.")
	md.AddHeading(2, "Subheading")
	md.AppendContent("Additional paragraph.")

	err := md.WriteFile("example.md")
	if err != nil {
		log.Fatalf("Error writing file: %v", err)
	}
}
```