package utils

import (
	"os"
	"path/filepath"
)

// 一些和文件相关的通用方法

// GetAllFiles 获取给定目录下所有的文件，包括子目录中的文件
// 源码来自filepath.Walk，去除最后的排序步骤
func GetAllFiles(dir string) (files []string, err error) {
	fileOpen, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	fileInfos, err := fileOpen.Readdir(-1)
	fileOpen.Close()
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range fileInfos {
		fileName := filepath.Join(dir, fileInfo.Name())
		if fileInfo.IsDir() {
			// 是目录，需要递归遍历
			childFiles, err := GetAllFiles(fileName)
			if err != nil {
				return nil, err
			}
			files = append(files, childFiles...)
		} else {
			files = append(files, fileName)
		}
	}
	return files, nil
}