package snstools

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

// MockAWS represents the AWS session
type MockAWS struct {
	Instance *session.Session
	Topic    string
}

// MockNewSession create mock aws session
func MockNewSession() *session.Session {
	return &session.Session{}
}

// MockBuildInstance create aws session with stored ~/.aws credentials
func MockBuildInstance(sess *session.Session) MockAWS {
	// sess, err := session.NewSession(&aws.Config{Region: aws.String("eu-west-3")})
	return MockAWS{Instance: sess}
}

// Send push to aws SNS
func (s MockAWS) Send(sns *snsiface.SNSAPI, msg string) (string, error) {
	return msg, nil
}
