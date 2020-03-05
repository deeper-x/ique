package snstools

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/golang/mock/gomock"
)

func TestPushToSNS(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSNSAPI := NewMockSNSAPI(ctrl)

	msg := "ab28f744-0bb7-534a-b541-a19ec8d5c7f"
	topic := "arn:aws:sns:eu-west-3:777350386990:justopic"

	mockSNSAPI.EXPECT().
		Publish(&sns.PublishInput{Message: &msg, TopicArn: &topic}).
		Return(&sns.PublishOutput{MessageId: &msg}, nil)

	_, err := PushToSNS(mockSNSAPI, msg)

	if err != nil {
		t.Errorf("Error in pushing SNS: %v", err)
	}
}
