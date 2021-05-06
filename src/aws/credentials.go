package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/giustech/dumper/src/variable"
)

func ConnectAws() *session.Session {
	envi:=variable.GetEnvironments()
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(envi.Region),
			Credentials: credentials.NewStaticCredentials(
				envi.AccessKeyId,
				envi.SecretAccessKey,
				"",
			),
		})
	if err != nil {
		panic(err)
	}
	return sess
}

