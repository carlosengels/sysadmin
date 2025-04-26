# Super Simple Transfer

A simple tool for syncing files with AWS S3.

## Prerequisites

Before using this tool, ensure you have:

1. AWS CLI installed
   - Installation instructions: https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html

2. AWS credentials configured
   - Run `aws configure` to set up your credentials
   - Ensure your credentials have S3 access permissions

## Setup

1. Run the setup script:
   ```bash
   ./syncsetup.sh
   ```

2. The script will:
   - Create a `~/bin` directory if it doesn't exist
   - Install the sync tool as `bsync`
   - Add `~/bin` to your PATH

3. Restart your terminal or run `source ~/.bashrc` to apply PATH changes

## Usage

The `bsync` command will be available in your terminal after setup.

### Default Behavior

By default, the tool:
- Uses `s3://supersimpletransferbucket` as the S3 bucket
- Creates and uses `~/tmp_store` as the local storage directory
- Performs a bi-directional sync (both upload and download)
- Uses exact timestamps for accurate sync
- Runs silently (output redirected to /dev/null)

### Configuration

You can modify the following environment variables before running `bsync`:
- `TMP_STORE`: Local directory path (default: `~/tmp_store`)
- `S3_BUCKET`: S3 bucket URL (default: `s3://supersimpletransferbucket`)

Example:
```bash
export TMP_STORE="/path/to/your/directory"
export S3_BUCKET="s3://your-bucket-name"
bsync
```

Note: Make sure you have the necessary AWS S3 bucket permissions before using the tool. 