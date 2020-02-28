package filesys

// MockFileManager is the FileManager mock
type MockFileManager struct {
	Pwd string
}

// ReadFileContent read file content and return it
func (mfm MockFileManager) ReadFileContent(fileName string) (string, error) {
	return "mock file content", nil
}

// AddListener listen for file creation
func (mfm MockFileManager) AddListener() error {
	return nil
}
