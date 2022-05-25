package util

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// GetFileSize 获取文件大小
func GetFileSize(filename string) int64 {
	var result int64
	_ = filepath.Walk(filename, func(path string, f os.FileInfo, err error) error {
		result = f.Size()
		return nil
	})
	return result
}

// GetCurrentPath 获取当前路径，比如：E:/abc/data/test
func GetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
