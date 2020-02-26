package client

import (
	"testing"
)

func TestClientRun(t *testing.T) {
	mp := MockPitch{}

	err := Run(&mp, "queue", "hello")

	if err != nil {
		t.Errorf("TestRun error %s", err)
	}
}
