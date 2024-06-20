package biz

import (
	"os"
	"path/filepath"
)

// 针对单个文件夹的处理

func ProcessFolder(folderPath string) ([]string, error) {
	markdownFiles, err := scanMarkdownFiles(folderPath)
	if err != nil {
		return nil, err
	}
	var blogsMd5Set []string
	for _, markdownFile := range markdownFiles {
		md5Path, err := ProcessFile(markdownFile)
		if err != nil {
			return nil, err
		}
		blogsMd5Set = append(blogsMd5Set, md5Path)
	}
	return blogsMd5Set, nil
}

// scanMarkdownFiles 扫描指定目录下的所有Markdown文件，不包括子文件夹
func scanMarkdownFiles(dir string) ([]string, error) {
	var markdownFiles []string

	// 打开目录
	f, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// 读取目录内容
	files, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}

	// 遍历目录内容
	for _, file := range files {
		// 如果是文件且扩展名为.md，则添加到Markdown文件列表中
		if !file.IsDir() && filepath.Ext(file.Name()) == ".md" {
			markdownFiles = append(markdownFiles, filepath.Join(dir, file.Name()))
		}
	}

	return markdownFiles, nil
}
