package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/exp/constraints"
)

func ListDir(dir string) (filePaths []string, baseNames []string, err error) {
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("err !!!\n")
			return err // 遇到错误时返回
		}
		if !info.IsDir() { // 确保是文件
			filePaths = append(filePaths, path)                                     // 添加完整路径到切片
			baseName := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)) // 去除文件后缀
			baseNames = append(baseNames, baseName)                                 // 添加去除后缀的文件名到切片
		}
		return nil
	})

	if err != nil {
		return nil, nil, err // 遇到错误时返回
	}

	return filePaths, baseNames, nil
}

func IsInSlice[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

func Min[T constraints.Ordered](x, y T) T {
	if x > y {
		return y
	}
	return x
}
