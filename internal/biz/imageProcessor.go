package biz

import (
	"blog-processor/global"
	"blog-processor/internal/model"
	"blog-processor/pkg/utils/encoder"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type ImageProcessor struct {
	Filepath   string
	Content    string
	OutputPath string
	BlogImages []model.BlogImage
}

func (processor *ImageProcessor) ProcessImage() error {
	// 编码输出文件夹名称并组合成绝对路径
	fileRelativePath, err := filepath.Rel(global.BasicSetting.BlogDir, processor.Filepath)
	if err != nil {
		return err
	}
	processor.OutputPath = global.BasicSetting.OutputDir + "/" + encoder.EncodeMD5(fileRelativePath)
	// 创建输出文件夹
	if err = os.MkdirAll(processor.OutputPath, os.ModePerm); err != nil {
		return err
	}

	// 提取博客图片对象
	processor.ExtractImages()

	// 写入新路径
	err = processor.WriteToNewDir()
	if err != nil {
		return err
	}
	// 替换
	processor.ReplaceContent()

	return nil
}

func (processor *ImageProcessor) ExtractImages() {
	// 正则表达式匹配Markdown图片路径
	re := regexp.MustCompile(`!\[.*?\]\((.*?)\)`)
	matches := re.FindAllStringSubmatch(processor.Content, -1)

	var imagePaths []string
	for _, match := range matches {
		// 提取非在线源的图片路径
		if !isOnlineSource(match[1]) {
			imagePaths = append(imagePaths, match[1])
		}
	}

	for _, imagePath := range imagePaths {
		blogImages := model.BlogImage{
			RawPath:         imagePath,
			AbsPath:         filepath.Join(filepath.Dir(processor.Filepath), imagePath),
			FileExt:         filepath.Ext(imagePath),
			FileName:        filepath.Base(imagePath),
			EncodedFileName: encoder.EncodeMD5(filepath.Base(processor.Filepath) + "=" + imagePath),
		}
		processor.BlogImages = append(processor.BlogImages, blogImages)
	}
}

// isOnlineSource 判断路径是否为在线源
func isOnlineSource(path string) bool {
	return strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://")
}

func (processor *ImageProcessor) WriteToNewDir() error {
	for _, blogImage := range processor.BlogImages {
		// 获取原始图片真实路径
		input, err := os.Open(blogImage.AbsPath)
		if err != nil {
			return err
		}
		defer input.Close()

		// 构建输出文件
		outputPath := filepath.Join(processor.OutputPath, blogImage.EncodedFileName+blogImage.FileExt)

		// 检查文件是否存在
		if _, err := os.Stat(outputPath); err == nil {
			// 文件已存在，跳过写入
			fmt.Printf("File %s already exists, skipping...\n", outputPath)
			continue
		} else if !os.IsNotExist(err) {
			// 其他错误
			return err
		}

		output, err := os.Create(outputPath)
		if err != nil {
			return err
		}
		defer output.Close()

		// 复制文件
		if _, err = io.Copy(output, input); err != nil {
			return err
		}
	}
	return nil
}

func (processor *ImageProcessor) ReplaceContent() {
	for _, blogImage := range processor.BlogImages {
		newURL := global.BasicSetting.HttpBasePath +
			filepath.Base(processor.OutputPath) + "/" +
			blogImage.EncodedFileName +
			blogImage.FileExt
		processor.Content = strings.Replace(processor.Content, blogImage.RawPath, newURL, -1)
	}
}
