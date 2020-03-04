package snstools

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// MockAWS represents the AWS session
type MockAWS struct {
	Instance *session.Session
	Topic    string
}

// MockBuildInstance create aws session with stored ~/.aws credentials
func MockBuildInstance() (MockAWS, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("eu-west-2")})

	return MockAWS{Instance: sess}, err
}

// Send push to aws SNS
func (s MockAWS) Send(msg string) (string, error) {
	return msg, nil
}
