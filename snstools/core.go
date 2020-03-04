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
	Instance *session.Session
	Topic    string
}

// NewSession return AWS session
func NewSession() *session.Session {
	instance := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return instance
}

// BuildInstance AWS session with stored ~/.aws credentials
func BuildInstance(sess *session.Session) AWS {
	return AWS{Instance: sess}
}

// Send notification to aws SNS
func (s AWS) Send(svc *sns.SNS, msg string) (string, error) {
	s.Topic = configuration.AwsTopic

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
func PushToSNS(s AWS, msg string) (string, error) {
	sess := NewSession()
	awsObj := BuildInstance(sess)

	svc := sns.New(s.Instance)

	res, err := awsObj.Send(svc, msg)
	if err != nil {
		return res, err
	}

	return res, nil
}
