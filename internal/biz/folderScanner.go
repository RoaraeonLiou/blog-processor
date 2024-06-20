package biz

import (
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func ScanFolder(basePath string) ([]string, error) {
	var folders []string
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果是文件夹且不是根目录，则添加到文件夹列表中
		if info.IsDir() && path != basePath {
			folders = append(folders, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	
	sortFoldersByDepth(folders)
	return folders, nil
}

// sortFoldersByDepth 按深度排序文件夹，越深的文件夹越靠前
func sortFoldersByDepth(folders []string) {
	sort.Slice(folders, func(i, j int) bool {
		return depth(folders[i]) > depth(folders[j])
	})
}

// depth 计算文件夹路径的深度
func depth(path string) int {
	return strings.Count(path, string(os.PathSeparator))
}
