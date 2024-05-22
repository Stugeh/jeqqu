package utils

import (
	"fmt"
	"io/fs"
	"os"
	"sort"
	"time"
)

const TEMP_DIR = "./temp/"
const MAX_FILES = 20

func createTempFilePath() string {
	return TEMP_DIR + fmt.Sprint(time.Now().Unix()) + "_temp.json"
}

func SortFilesByModTimeAsc(fileInfos []fs.FileInfo) {
	sort.Slice(fileInfos, func(i, j int) bool {
		return fileInfos[i].ModTime().Before(fileInfos[j].ModTime())
	})
}

func SortFilesByModTimeDesc(fileInfos []fs.FileInfo) {
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
		SortFilesByModTimeAsc(fileInfos)
		os.Remove(TEMP_DIR + fileInfos[0].Name())
	}

	path := createTempFilePath()
	err = os.WriteFile(path, bytes, 0644)

	return path, err
}
