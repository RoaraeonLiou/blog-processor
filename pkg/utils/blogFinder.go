package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func FindAllMarkdownFiles(dirPath string) ([]string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fmt.Println(wd)
	var files []string
	err = filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".md" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}
