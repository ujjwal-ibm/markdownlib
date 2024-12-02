package markdownlib

import (
	"errors"
	"strings"
)

// ValidateHeadingLevels ensures all heading levels are between 1 and 6.
func (mf *MarkdownFile) ValidateHeadingLevels() error {
	lines := strings.Split(mf.Content, "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "#") {
			level := strings.Count(strings.Split(line, " ")[0], "#")
			if level < 1 || level > 6 {
				return errors.New("invalid heading level detected")
			}
		}
	}
	return nil
}

// ValidateStructure checks for at least one heading in the document.
func (mf *MarkdownFile) ValidateStructure() error {
	headings, _ := mf.ParseHeadings()
	if len(headings) == 0 {
		return errors.New("document has no headings")
	}
	return nil
}
