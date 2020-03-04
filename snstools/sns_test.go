package snstools

import (
	"testing"
)

var awsObj = MockAWS{}

func TestPushToSNS(t *testing.T) {
	mockClient := NewClient()

	output, err := PushToSNS(mockClient, "demo from test")
	if err != nil {
		t.Errorf("Error in sending %v to SNS: %v", output, err)
	}
}
