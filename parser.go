package markdownlib

import (
	"errors"
	"strings"
)

// ParseLinks extracts all links in the format [text](url) from the content.
func (mf *MarkdownFile) ParseLinks() ([]string, error) {
	lines := strings.Split(mf.Content, "\n")
	var links []string

	for _, line := range lines {
		start := strings.Index(line, "[")
		end := strings.Index(line, "](")
		if start != -1 && end != -1 {
			link := line[start+1 : end]
			links = append(links, link)
		}
	}

	if len(links) == 0 {
		return nil, errors.New("no links found")
	}
	return links, nil
}

// ParseLists extracts all list items from the Markdown content.
func (mf *MarkdownFile) ParseLists() []string {
	lines := strings.Split(mf.Content, "\n")
	var lists []string

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "- ") || strings.HasPrefix(trimmed, "* ") {
			lists = append(lists, trimmed)
		}
	}
	return lists
}
