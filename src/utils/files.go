package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"time"
)

const TEMP_DIR = "./temp/"
const MAX_FILES = 20

func createTempFilePath() string {
	return TEMP_DIR + fmt.Sprint(time.Now().Unix()) + "_temp.json"
}

func GetNewestTempFilePath() (string, error) {
	dir, err := os.Open(TEMP_DIR)

	if err != nil {
		return "", err
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return "", err
	}

	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].ModTime().After(fileInfos[j].ModTime())
	})

	if len(fileInfos) > 0 {
		newestFilePath := filepath.Join(TEMP_DIR, fileInfos[0].Name())
		return newestFilePath, nil
	}

	return "", fmt.Errorf("no files found in directory")
}

func SortFileInfoAsc(fileInfos []fs.FileInfo) {
	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].ModTime().Before(fileInfos[j].ModTime())
	})
}

func NewTempFile(bytes []byte) (string, error) {
	dir, err := os.Open(TEMP_DIR)
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return "", err
	}

	// Delete oldest temp file if Max reached
	if len(fileInfos) > MAX_FILES {
		SortFileInfoAsc(fileInfos)
		os.Remove(TEMP_DIR + fileInfos[0].Name())
	}

	path := createTempFilePath()
	err = os.WriteFile(path, bytes, 0644)

	return path, err
}
