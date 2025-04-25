#!/bin/bash

# Define environment variable if not already set
TMP_STORE="${TMP_STORE:-/home/engels/temp_store}"
S3_BUCKET="s3://supersimpletransferbucket"

# Create local directory if it doesn't exist
mkdir -p "$TMP_STORE"

aws s3 sync "$TMP_STORE" "$S3_BUCKET" --exact-timestamps

aws s3 sync "$S3_BUCKET" "$TMP_STORE" --exact-timestamps

echo "Sync complete."ho