package app

import (
	"os"
	"path/filepath"
	"strings"
)

// FileExists 判断文件是否存在
func FileExists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}

// FilePut 将数据存入文件
func FilePut(data []byte, to string) error {
	err := os.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// FileNameWithoutExtension 去除后缀名
func FileNameWithoutExtension(fileName string) string {
	return strings.TrimSuffix(fileName, filepath.Ext(fileName))
}
