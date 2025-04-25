#!/bin/bash

# Get the user's home directory
USER_HOME="${HOME}"

# Define environment variable if not already set
TMP_STORE="$USER_HOME/tmp_store"
S3_BUCKET="s3://supersimpletransferbucket"

# Create local directory if it doesn't exist
mkdir -p "$TMP_STORE"

aws s3 sync "$TMP_STORE" "$S3_BUCKET" --exact-timestamps > /dev/null 2>&1

aws s3 sync "$S3_BUCKET" "$TMP_STORE" --exact-timestamps > /dev/null 2>&1

echo "Sync complete."