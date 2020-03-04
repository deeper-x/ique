package snstools

import "testing"

var awsObj = MockAWS{}

func TestPushToSNS(t *testing.T) {
	awsObj := MockBuildInstance()
	output, err := PushToSNS(awsObj, "demo from test")

	if err != nil {
		t.Errorf("Error in sending %v to SNS: %v", output, err)
	}
}
