package utils

import (
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"time"
)

func GetDownloadDir() (string, error) {
	userCacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	downloadDir := path.Join(userCacheDir, "hypotd")
	if err := os.MkdirAll(downloadDir, 0755); err != nil {
		return "", err
	}

	return downloadDir, nil
}

func ClearOldFiles(directory string, daysOlderThan uint32) error {
	olderThan := time.Now().Add(-time.Duration(daysOlderThan*24) * time.Hour)

	return filepath.Walk(directory, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if info.ModTime().After(olderThan) {
			return nil
		}

		return os.Remove(path)
	})
}
