#!/bin/bash

# Get the user's home directory
USER_HOME="${HOME}"

# Define environment variable if not already set
TMP_STORE="$USER_HOME/tmp_store"
S3_BUCKET="s3://supersimpletransferbucket"

# Create local directory if it doesn't exist
mkdir -p "$TMP_STORE"

# Define log file
LOG_FILE="$TMP_STORE/sync.log"

# Get host information
HOSTNAME=$(hostname)
TIMESTAMP=$(date '+%Y-%m-%d %H:%M:%S')

# Log sync start
echo "[$TIMESTAMP] [$HOSTNAME] Starting sync operation" >> "$LOG_FILE"

# Sync local to S3
echo "[$TIMESTAMP] [$HOSTNAME] Syncing local to S3: $TMP_STORE -> $S3_BUCKET" >> "$LOG_FILE"
aws s3 sync "$TMP_STORE" "$S3_BUCKET" --exact-timestamps >> "$LOG_FILE" 2>&1

# Sync S3 to local
echo "[$TIMESTAMP] [$HOSTNAME] Syncing S3 to local: $S3_BUCKET -> $TMP_STORE" >> "$LOG_FILE"
aws s3 sync "$S3_BUCKET" "$TMP_STORE" --exact-timestamps >> "$LOG_FILE" 2>&1

# Log sync completion
echo "[$TIMESTAMP] [$HOSTNAME] Sync operation completed" >> "$LOG_FILE"
echo "Sync complete. Check $LOG_FILE for details."