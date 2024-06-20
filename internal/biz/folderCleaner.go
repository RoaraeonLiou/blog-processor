package biz

import (
	"fmt"
	"os"
	"path/filepath"
)

func DeleteEmptyDir(dirPath string) error {
	isEmpty, err := isDirEmpty(dirPath)
	if err != nil {
		return err
	}

	if isEmpty {
		err = deleteHiddenFiles(dirPath)
		if err != nil {
			return err
		}

		if err := os.Remove(dirPath); err != nil {
			return err
		}
		fmt.Printf("Deleted empty directory: %s\n", dirPath)
	}

	return nil
}

func isDirEmpty(dirPath string) (bool, error) {
	f, err := os.Open(dirPath)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// 读取目录内容
	names, err := f.Readdirnames(-1)
	if err != nil {
		return false, err
	}

	// 检查是否只有隐藏文件
	for _, name := range names {
		if name[0] != '.' {
			return false, nil
		}
	}

	// 目录为空或只有隐藏文件
	return true, nil
}

func deleteHiddenFiles(dirPath string) error {
	f, err := os.Open(dirPath)
	if err != nil {
		return err
	}
	defer f.Close()

	// 读取目录内容
	names, err := f.Readdirnames(-1)
	if err != nil {
		return err
	}

	// 删除所有隐藏文件
	for _, name := range names {
		if name[0] == '.' {
			hiddenFilePath := filepath.Join(dirPath, name)
			if err := os.Remove(hiddenFilePath); err != nil {
				return err
			}
			fmt.Printf("Deleted hidden file: %s\n", hiddenFilePath)
		}
	}

	return nil
}
