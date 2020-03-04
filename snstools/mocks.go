package snstools

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

// MockAWS represents the AWS session
type MockAWS struct {
	Instance *session.Session
	Topic    string
}

// MockBuildInstance create aws session with stored ~/.aws credentials
func MockBuildInstance() MockAWS {
	sess := &session.Session{}

	return MockAWS{Instance: sess}
}

// Send push to aws SNS
func (s MockAWS) Send(msg string) (string, error) {
	return msg, nil
}
