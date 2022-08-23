package pkg

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"log"
)

func newAws() (aws.Config, error) {

	region := "eu-central-1"
	profile := "default"
	awsClient, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(region),
		config.WithSharedConfigProfile(profile),
	)
	if err != nil {
		return aws.Config{}, fmt.Errorf("unable to initialize AWS service, error: %v", err)
	}

	return awsClient, nil
}

func CreateBucket(dryRun bool, bucketName string) error {

	log.Println("createBucketCalled")
	if dryRun {
		log.Printf("[#99] Dry-run mode, bucket creation skipped:  %s", bucketName)
		return nil
	}

	// todo: use method approach to avoid new AWS client initializations
	awsClient, err := newAws()
	if err != nil {
		log.Printf("failed to attempt bucket creation, error: %v ", err)
		return err
	}

	s3Client := s3.NewFromConfig(awsClient)

	log.Println("creating bucket: ", bucketName)

	regionName := "eu-central-1"
	log.Println("region is ", regionName)

	_, err = s3Client.CreateBucket(
		context.Background(),
		&s3.CreateBucketInput{
			Bucket: &bucketName,
			CreateBucketConfiguration: &s3Types.CreateBucketConfiguration{
				LocationConstraint: s3Types.BucketLocationConstraint(regionName),
			},
		})
	if err != nil {
		return err
	}

	log.Printf("bucket.%s.created", bucketName)

	return nil
}
