package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	AWS struct {
		Region string
		Bucket string
	}
	Directories struct {
		Upload   string
		Download string
	}
	Sync struct {
		DeleteMissing bool
	}
}

func Load() *Config {
	config := &Config{}
	
	// AWS settings
	config.AWS.Region = getEnvOrDefault("AWS_REGION", "us-east-1")
	config.AWS.Bucket = getEnvOrDefault("AWS_BUCKET", "")
	
	// Directory settings
	config.Directories.Upload = getEnvOrDefault("UPLOAD_DIR", "./upload")
	config.Directories.Download = getEnvOrDefault("DOWNLOAD_DIR", "./download")
	
	// Sync settings
	config.Sync.DeleteMissing = getEnvOrDefault("DELETE_MISSING", "false") == "true"

	// Make paths absolute
	exePath, err := os.Executable()
	if err == nil {
		exeDir := filepath.Dir(exePath)
		config.Directories.Upload = filepath.Join(exeDir, config.Directories.Upload)
		config.Directories.Download = filepath.Join(exeDir, config.Directories.Download)
	}
	
	return config
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
} 