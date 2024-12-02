package markdownlib

import (
	"os"
	"testing"
)

func TestReadWrite(t *testing.T) {
	testPath := "test.md"
	testContent := "# Test Heading\nThis is a test."

	// Write test
	err := WriteMarkdown(testPath, testContent)
	if err != nil {
		t.Fatalf("Failed to write markdown file: %v", err)
	}

	// Read test
	readContent, err := ReadMarkdown(testPath)
	if err != nil {
		t.Fatalf("Failed to read markdown file: %v", err)
	}

	if readContent != testContent {
		t.Fatalf("Content mismatch: got %v, want %v", readContent, testContent)
	}

	// Cleanup
	_ = os.Remove(testPath)
}
