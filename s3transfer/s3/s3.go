package s3

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	client *s3.Client
	bucket string
}

func NewClient(region, bucket string) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config: %v", err)
	}

	return &Client{
		client: s3.NewFromConfig(cfg),
		bucket: bucket,
	}, nil
}

func (c *Client) UploadDir(dir string) error {
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		// Get relative path for S3 key
		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}

		// Open the file
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Upload to S3
		_, err = c.client.PutObject(context.TODO(), &s3.PutObjectInput{
			Bucket: aws.String(c.bucket),
			Key:    aws.String(relPath),
			Body:   file,
		})
		if err != nil {
			return fmt.Errorf("failed to upload %s: %v", path, err)
		}

		fmt.Printf("Uploaded: %s\n", relPath)
		return nil
	})
}

func (c *Client) DownloadDir(dir string) error {
	// List objects in bucket
	listInput := &s3.ListObjectsV2Input{
		Bucket: aws.String(c.bucket),
	}
	paginator := s3.NewListObjectsV2Paginator(c.client, listInput)

	// Process each object
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			return fmt.Errorf("failed to list objects: %v", err)
		}

		for _, obj := range page.Contents {
			key := *obj.Key
			localPath := filepath.Join(dir, key)

			// Create directory if it doesn't exist
			if err := os.MkdirAll(filepath.Dir(localPath), 0755); err != nil {
				return fmt.Errorf("failed to create directory for %s: %v", key, err)
			}

			// Download the object
			result, err := c.client.GetObject(context.TODO(), &s3.GetObjectInput{
				Bucket: aws.String(c.bucket),
				Key:    aws.String(key),
			})
			if err != nil {
				return fmt.Errorf("failed to get object %s: %v", key, err)
			}
			defer result.Body.Close()

			// Create the file
			file, err := os.Create(localPath)
			if err != nil {
				return fmt.Errorf("failed to create file %s: %v", localPath, err)
			}
			defer file.Close()

			// Copy the object content to the file
			if _, err := file.ReadFrom(result.Body); err != nil {
				return fmt.Errorf("failed to write file %s: %v", localPath, err)
			}

			fmt.Printf("Downloaded: %s\n", key)
		}
	}

	return nil
} 