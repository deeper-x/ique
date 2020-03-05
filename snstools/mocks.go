package snstools

import (
	"log"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/deeper-x/ique/configuration"
)

// MockAWS Define a mock struct to be used in your unit tests of myFunc.
type MockAWS struct {
	Topic  string
	Client snsiface.SNSAPI
}

// NewMockAWS return AWS session
func NewMockAWS(c snsiface.SNSAPI) *MockAWS {
	return &MockAWS{Client: c}
}

// Send notification to aws SNS
func (m *MockAWS) Send(msg string) (string, error) {
	m.Topic = configuration.AwsTopic

	result, err := m.Client.Publish(&sns.PublishInput{
		Message:  &msg,
		TopicArn: &m.Topic,
	})

	if err != nil {
		log.Println(err.Error())
		return msg, err
	}

	log.Println(*result.MessageId)

	return msg, nil
}

// Send notification to aws SNS
// func (s *AWS) Send(msg string) (string, error) {
// 	s.Topic = configuration.AwsTopic

// 	result, err := s.Client.Publish(&sns.PublishInput{
// 		Message:  &msg,
// 		TopicArn: &s.Topic,
// 	})

// 	if err != nil {
// 		log.Println(err.Error())
// 		return msg, err
// 	}

// 	log.Println(*result.MessageId)

// 	return msg, nil
// }
