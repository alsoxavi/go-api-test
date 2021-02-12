package awssdk

import (
	"log"

	enviroment "github.com/alsoxavi/go-api-test/config"
	"github.com/alsoxavi/go-api-test/model"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	session "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// SendSMS godoc
func SendSMS(sms *model.SMS) (string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(enviroment.Config("AWS_REGION")),
		Credentials: credentials.NewStaticCredentials(
			enviroment.Config("AWS_ACCESS_KEY_ID"),
			enviroment.Config("AWS_ACCESS_KEY"),
			"",
		),
	})
	if err != nil {
		log.Println(err)
		return "", err
	}

	client := sns.New(sess)
	input := &sns.PublishInput{
		Message:     aws.String(sms.Message),
		PhoneNumber: aws.String(sms.PhoneNumber),
	}

	result, err := client.Publish(input)
	if err != nil {
		return "", err
	}

	return result.String(), nil
}
