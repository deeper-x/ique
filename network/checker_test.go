package network

import "testing"

func TestServiceIsRunning(t *testing.T) {
	ms := MockService{URI: "192.168.1.1:80"}

	is := ms.IsRunning()

	if !is {
		t.Errorf("host %v should be available", ms.URI)
	}
}
