package snstools

import (
	"log"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/deeper-x/ique/configuration"
)

// MockAWS represents the AWS data
type MockAWS struct {
	Topic  string
	Client snsiface.SNSAPI
}

// Define a mock struct to be used in your unit tests of myFunc.
type mockSNSClient struct {
	snsiface.SNSAPI
}

func (m *mockSNSClient) AddPermission(input *sns.AddPermissionInput) (*sns.AddPermissionOutput, error) {
	// mock response/functionality
	return &sns.AddPermissionOutput{}, nil
}

// Send notification to aws SNS
func (s MockAWS) Send(msg string) (string, error) {
	s.Topic = configuration.AwsTopic

	result, err := s.Client.Publish(&sns.PublishInput{
		Message:  &msg,
		TopicArn: &s.Topic,
	})

	if err != nil {
		log.Println(err.Error())
		return msg, err
	}

	log.Println(*result.MessageId)

	return msg, nil
}
