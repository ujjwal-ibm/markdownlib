package markdownlib

import (
	"errors"
	"io/ioutil"
	"strings"
)

// MarkdownFile represents a Markdown document.
type MarkdownFile struct {
	Content string
}

// NewMarkdownFile creates a new Markdown file instance.
func NewMarkdownFile(content string) *MarkdownFile {
	return &MarkdownFile{Content: content}
}

// ReadFile reads a Markdown file from disk.
func ReadFile(path string) (*MarkdownFile, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewMarkdownFile(string(data)), nil
}

// WriteFile writes the Markdown content to disk.
func (mf *MarkdownFile) WriteFile(path string) error {
	return ioutil.WriteFile(path, []byte(mf.Content), 0644)
}

// AppendContent appends content to the Markdown file.
func (mf *MarkdownFile) AppendContent(content string) {
	mf.Content += "\n" + content
}

// ParseHeadings extracts all headings from the Markdown content.
func (mf *MarkdownFile) ParseHeadings() ([]string, error) {
	lines := strings.Split(mf.Content, "\n")
	var headings []string
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			headings = append(headings, line)
		}
	}
	return headings, nil
}

// AddHeading adds a heading to the Markdown content.
func (mf *MarkdownFile) AddHeading(level int, text string) error {
	if level < 1 || level > 6 {
		return errors.New("heading level must be between 1 and 6")
	}
	mf.AppendContent(strings.Repeat("#", level) + " " + text)
	return nil
}
