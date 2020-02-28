package filesys

import (
	"testing"
)

// AddListener(string) error
// ReadFileContent(string) (string, error)

func TestAddListener(t *testing.T) {
	mockFilesys := MockFileManager{}

	err := RunListen(mockFilesys)

	if err != nil {
		t.Errorf("Error running listener: %v", err)
	}
}
