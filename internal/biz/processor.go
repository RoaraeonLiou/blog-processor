package biz

import (
	"blog-processor/pkg/utils"
	"os"
	"path/filepath"
)

func ProcessSingleFile(filePath string, httpBasePath string, imgOutPath string) error {
	header, fileContent, err := utils.ReadYAMLHeader(filePath)
	if err != nil {
		return err
	}
	// todo: 处理头部

	// 创建文件夹
	fileName := filepath.Base(filePath)
	dirPath := filepath.Dir(filePath)
	folderName := filepath.Base(dirPath)
	imageFolderName := utils.EncodeMD5(fileName)
	fatherFolderName := utils.EncodeMD5(folderName)
	newImageDirPath := imgOutPath + "/" + fatherFolderName + "/" + imageFolderName
	if err := os.MkdirAll(newImageDirPath, os.ModePerm); err != nil {
		return err
	}

	imagesPaths := utils.ExtractImagesAndEncodeFilename(fileContent)

	newContent := utils.ReplaceImagePaths(fileContent, imagesPaths, httpBasePath, newImageDirPath)

	err = utils.CopyImagesToDir(imagesPaths, newImageDirPath, dirPath)
	if err != nil {
		return err
	}
	err = utils.WriteProcessedMarkdown(filePath, header, newContent)
	if err != nil {
		return err
	}
	return nil
}

func ProcessAll(filePath string, httpBasePath string, imgOutPath string) error {
	mdFilePaths, err := utils.FindAllMarkdownFiles(filePath)
	if err != nil {
		return err
	}
	for _, filePath := range mdFilePaths {
		err = ProcessSingleFile(filePath, httpBasePath, imgOutPath)
		if err != nil {
			return err
		}
	}
	return nil
}
