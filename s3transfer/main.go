package main

import (
	"fmt"
	"os"

	"github.com/yourusername/s3transfer/config"
	"github.com/yourusername/s3transfer/s3"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Load configuration
	cfg := config.Load()
	if cfg.AWS.Bucket == "" {
		fmt.Println("Error: AWS_BUCKET environment variable must be set")
		os.Exit(1)
	}

	// Create S3 client
	client, err := s3.NewClient(cfg.AWS.Region, cfg.AWS.Bucket)
	if err != nil {
		fmt.Printf("Error creating S3 client: %v\n", err)
		os.Exit(1)
	}

	command := os.Args[1]
	switch command {
	case "up":
		if err := upload(client, cfg); err != nil {
			fmt.Printf("Error uploading: %v\n", err)
			os.Exit(1)
		}
	case "down":
		if err := download(client, cfg); err != nil {
			fmt.Printf("Error downloading: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Usage: s3transfer <command>")
	fmt.Println("\nCommands:")
	fmt.Println("  up    Upload files from upload/ to S3")
	fmt.Println("  down  Download files from S3 to download/")
	fmt.Println("\nConfiguration (via environment variables):")
	fmt.Println("  AWS_REGION      AWS region (default: us-east-1)")
	fmt.Println("  AWS_BUCKET      S3 bucket name (required)")
	fmt.Println("  UPLOAD_DIR      Upload directory (default: ./upload)")
	fmt.Println("  DOWNLOAD_DIR    Download directory (default: ./download)")
	fmt.Println("  DELETE_MISSING  Delete files not in source (default: false)")
}

func upload(client *s3.Client, cfg *config.Config) error {
	fmt.Printf("Uploading files from %s to s3://%s...\n", 
		cfg.Directories.Upload, cfg.AWS.Bucket)
	return client.UploadDir(cfg.Directories.Upload)
}

func download(client *s3.Client, cfg *config.Config) error {
	fmt.Printf("Downloading files from s3://%s to %s...\n", 
		cfg.AWS.Bucket, cfg.Directories.Download)
	return client.DownloadDir(cfg.Directories.Download)
} 