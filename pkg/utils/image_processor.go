package utils

import (
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// 提取图片地址并返回MD5编码的文件名
func ExtractImagesAndEncodeFilename(content string) []string {
	re := regexp.MustCompile(`!\[.*?\]\((.*?)\)`) // 正则表达式匹配Markdown图片路径
	matches := re.FindAllStringSubmatch(content, -1)

	var imagePaths []string
	for _, match := range matches {
		if !strings.HasPrefix(match[1], "http") {
			imagePaths = append(imagePaths, match[1])
		}
	}
	return imagePaths
}

// 替换文件中的图片路径为新的HTTP路径，仅针对非HTTP源的图片
func ReplaceImagePaths(content string, imagePaths []string, httpBasePath string, newImageDirPath string) string {
	for _, imagePath := range imagePaths {
		if !strings.HasPrefix(imagePath, "http") {
			md5Filename := EncodeMD5(filepath.Base(imagePath)) + filepath.Ext(imagePath)
			newURL := httpBasePath + newImageDirPath + "/" + md5Filename
			content = strings.Replace(content, imagePath, newURL, -1)
		}
	}
	return content
}

// 复制文件到新目录并重命名为MD5编码后的名称
func CopyImagesToDir(imagePaths []string, dirPath string, imgSourcePath string) error {
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return err
	}

	for _, imagePath := range imagePaths {
		realImagePath := imgSourcePath + "/" + imagePath
		input, err := os.Open(realImagePath)
		if err != nil {
			return err
		}
		defer input.Close()

		outputPath := filepath.Join(dirPath, EncodeMD5(filepath.Base(imagePath))+filepath.Ext(imagePath))
		output, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		defer output.Close()

		if _, err = io.Copy(output, input); err != nil {
			return err
		}
	}
	return nil
}
