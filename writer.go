package markdownlib

import (
	"os"
)

// WriteMarkdown writes the given content to a Markdown file.
func WriteMarkdown(path, content string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return err
}
