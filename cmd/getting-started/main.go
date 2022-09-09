package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

var (
	MyTopic  = aws.String("go-topic")
	MyBucket = aws.String("go-bucket")
)

func main() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	// Create an Amazon S3 service client
	s3Client := s3.NewFromConfig(cfg)
	// Get the first page of results for ListObjectsV2 for a bucket
	GetBucketExample(s3Client)

	// Create an Amazon SNS service client
	snsClient := sns.NewFromConfig(cfg)
	// Get the topics
	CreateTopicExample(snsClient)
}

func GetBucketExample(client *s3.Client) {
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: MyBucket,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("first page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}
}

func CreateTopicExample(client *sns.Client) {
	parms := &sns.CreateTopicInput{
		Name: MyTopic,
	}

	results, err := client.CreateTopic(context.TODO(), parms)
	if err != nil {
		panic("sns error, " + err.Error())
	}

	fmt.Println(*results.TopicArn)
}
