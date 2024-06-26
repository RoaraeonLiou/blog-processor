package biz

import (
	"blog-processor/global"
	"blog-processor/internal/biz/processor"
	"blog-processor/internal/model"
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
		// 如果是文件夹且不是根目录且不是隐藏文件夹，则添加到文件夹列表中
		if info.IsDir() && path != basePath && !strings.HasPrefix(info.Name(), ".") {
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

func ScanFolderWithCommonHeader(basePath string) ([]string, map[string]*model.BlogHeader, error) {
	var headerMap = make(map[string]*model.BlogHeader)
	folders, err := ScanFolder(basePath)
	if err != nil {
		return nil, nil, err
	}
	for _, folder := range folders {
		newHeader, err := findCommonHeader(folder)
		if err != nil {
			return nil, nil, err
		}
		headerMap[folder] = newHeader
	}
	return folders, headerMap, nil
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

func findCommonHeader(dirPath string) (*model.BlogHeader, error) {
	hiddenFilePath := filepath.Join(dirPath, global.BasicSetting.CommonHeaderFileName+global.BasicSetting.CommonHeaderFileExt)
	// 检查文件是否存在且是隐藏文件
	_, err := os.Stat(hiddenFilePath)
	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	var commonHeader *model.BlogHeader
	if err == nil {
		headerFileContent, err := os.ReadFile(hiddenFilePath)
		if err != nil {
			return nil, err
		}
		commonHeader, err = processor.ParseHeader(string(headerFileContent), global.BasicSetting.CommonHeaderFileExt, global.BasicSetting.CommonHeaderFileFormat)
		if err != nil {
			return nil, err
		}
	}
	return commonHeader, nil
}
