// Package snstools is the AWS::SNS core component
package snstools

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
	"github.com/deeper-x/ique/configuration"
)

// Sender is the aws session main interface
type Sender interface {
	Send(snsiface.SNSAPI, string) (string, error)
}

// AWS represents the AWS data
type AWS struct {
	Topic  string
	Client snsiface.SNSAPI
}

// NewClient return AWS session
func NewClient() snsiface.SNSAPI {
	instance := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return sns.New(instance)
}

// Send notification to aws SNS
func (s *AWS) Send(msg string) (string, error) {
	s.Topic = configuration.AwsTopic

	result, err := s.Client.Publish(&sns.PublishInput{
		Message:  &msg,
		TopicArn: &s.Topic,
	})

	if err != nil {
		log.Println(err.Error())
		return msg, err
	}

	log.Println(*result)

	return msg, nil
}

// PushToSNS implements the sns pushing
func PushToSNS(client snsiface.SNSAPI, msg string) (string, error) {
	a := AWS{Client: client}

	res, err := a.Send(msg)
	if err != nil {
		return res, err
	}

	return res, nil
}
