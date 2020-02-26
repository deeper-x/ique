package server

import "testing"

func TestRunServer(t *testing.T) {
	ma := MockAgent{}

	err := Run(&ma, "demo")

	if err != nil {
		t.Errorf("RunServer test failed: %v", err)
	}
}
