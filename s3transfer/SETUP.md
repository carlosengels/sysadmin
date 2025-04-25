# S3 Transfer Tool Setup Guide

## Progress
✅ AWS Credentials Setup
✅ Directory Setup
✅ S3 Bucket Configuration (CloudFormation)
✅ Configuration Setup
✅ S3 Operations
✅ Sync Engine

## Overview
The S3 Transfer Tool is a simple command-line utility for moving files between your computer and Amazon S3. It provides two basic commands: upload and download.

## 1. AWS Credentials Setup

### Prerequisites
- AWS Account
- AWS CLI installed

### Steps to Configure AWS CLI Credentials
```bash
aws configure
```
- Enter your AWS Access Key ID
- Enter your AWS Secret Access Key
- Enter your default region (e.g., us-east-1)
- Enter your default output format (json)

## 2. Directory Setup

### Simple Directory Structure
```
s3transfer/
├── upload/    # Put files here to upload to S3
└── download/  # Files from S3 will appear here
```

### Directory Creation
```bash
mkdir -p s3transfer/{upload,download}
```

## 3. S3 Bucket Configuration

The S3 bucket configuration is defined in the CloudFormation template located at `cloudformation/s3bucket.yaml`. This template sets up a secure S3 bucket with encryption.

### Deploy CloudFormation Stack
```bash
aws cloudformation create-stack \
  --stack-name s3transfer-bucket \
  --template-body file://cloudformation/s3bucket.yaml \
  --capabilities CAPABILITY_IAM
```

## 4. Configuration

The tool is configured using environment variables:

```bash
# Required
export AWS_BUCKET="your-bucket-name"

# Optional (with defaults)
export AWS_REGION="us-east-1"
export UPLOAD_DIR="./upload"
export DOWNLOAD_DIR="./download"
export DELETE_MISSING="false"
```

## 5. Usage Examples

### Upload Files to S3
```bash
# Set up environment
export AWS_BUCKET="my-s3-bucket"
export AWS_REGION="us-east-1"

# Copy files to upload directory
cp my-files/* s3transfer/upload/

# Upload to S3
s3transfer up
```

### Download Files from S3
```bash
# Set up environment
export AWS_BUCKET="my-s3-bucket"
export AWS_REGION="us-east-1"

# Download from S3
s3transfer down

# Files will appear in download directory
ls s3transfer/download/
```

## Implementation Details

### Features
- Simple two-command interface (`up` and `down`)
- Preserves directory structure
- Handles large files automatically
- Progress feedback for each file
- Environment-based configuration

### Error Handling
- Validates AWS credentials
- Checks for required bucket name
- Reports file-specific errors
- Graceful error handling for network issues

### Security
- Uses AWS SDK v2 for secure communication
- Relies on AWS credentials management
- Server-side encryption in S3 bucket
- No local credential storage

## Next Steps
1. Test the implementation
2. Create a release build
3. Add to package managers (optional) 