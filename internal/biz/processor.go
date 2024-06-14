package biz

import "blog-processor/pkg/utils"

func ProcessSingleFile(filePath string, httpBasePath string) error {
	header, fileName, fileContent, err := utils.ReadYAMLHeader(filePath)
	if err != nil {
		return err
	}
	// todo: 处理头部

	imagesPath, newImageDir := ExtractImagesAndEncodeFilename(fileContent, fileName)

	newContent := ReplaceImagePaths(fileContent, newImageDir, httpBasePath)

	err = CopyImagesToDir(imagesPath, newImageDir)
	if err != nil {
		return err
	}
	err = WriteProcessedMarkdown(filePath, header, newContent)
	if err != nil {
		return err
	}
	return nil
}

func ProcessAll(filePath string, httpBasePath string) error {
	mdFilePaths, err := FindAllMarkdownFiles(filePath)
	if err != nil {
		return err
	}
	for _, filePath := range mdFilePaths {
		err = ProcessSingleFile(filePath, httpBasePath)
		if err != nil {
			return err
		}
	}
	return nil
}
