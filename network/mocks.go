package network

// MockService is the core instance
type MockService struct {
	URI string
}

// IsRunning check if instance is running
func (ms MockService) IsRunning() bool {
	return true
}
