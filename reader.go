package markdownlib

import (
	"io/ioutil"
)

// ReadMarkdown reads the entire content of a Markdown file.
func ReadMarkdown(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
