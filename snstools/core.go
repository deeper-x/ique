package snstools

import (
	"log"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/deeper-x/ique/configuration"
)

// Sender is the aws session main interface
type Sender interface {
	Send(string) (string, error)
}

// AWS represents the AWS data
type AWS struct {
	Instance *session.Session
	Topic    string
}

// BuildInstance AWS session with stored ~/.aws credentials
func BuildInstance() AWS {
	instance := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return AWS{Instance: instance}
}

// Send notification to aws SNS
func (s AWS) Send(msg string) (string, error) {
	s.Topic = configuration.AwsTopic

	svc := sns.New(s.Instance)

	result, err := svc.Publish(&sns.PublishInput{
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

// PushToSNS implements the sns pushing
func PushToSNS(s Sender, msg string) (string, error) {
	awsObj := BuildInstance()

	res, err := awsObj.Send(msg)
	if err != nil {
		return res, err
	}

	return res, nil
}
