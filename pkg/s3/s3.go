package s3

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Client struct {
	client       *s3.S3
	sess         *session.Session
	customHost   string
	removeBucket bool
}

func NewClient(accessKey, secretKey string, options ...Options) (*Client, error) {
	client := &Client{}
	config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
		Region:           aws.String("ams1"),
		S3ForcePathStyle: aws.Bool(false),
	}
	if len(options) > 0 {
		option := options[0]
		client.removeBucket = option.RemoveBucket
		if option.EndPoint != "" {
			config.Endpoint = aws.String(option.EndPoint)
		}
		if option.Region != "" {
			config.Region = aws.String(option.Region)
		}
		if option.CustomHost != "" {
			client.customHost = option.CustomHost
		}
		if option.S3ForcePathStyle != nil {
			config.S3ForcePathStyle = option.S3ForcePathStyle
		}
	}
	sess, err := session.NewSession(config)
	if err != nil {
		return nil, err
	}
	client.sess = sess
	client.client = s3.New(sess)
	return client, nil
}
