package utils

import (
	"blog-processor/internal/model"
	"bufio"
	"gopkg.in/yaml.v3"
	"os"
)

func WriteProcessedMarkdown(filePath string, header *model.BlogHeader, bodyContent string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString("---\n")
	if err != nil {
		return err
	}

	headerYAML, err := yaml.Marshal(header)
	if err != nil {
		return err
	}

	_, err = writer.Write(headerYAML)
	if err != nil {
		return err
	}
	_, err = writer.WriteString("---\n")
	if err != nil {
		return err
	}
	_, err = writer.WriteString(bodyContent)
	if err != nil {
		return err
	}

	err = writer.Flush()
	if err != nil {
		return err
	}

	return nil
}

func ReadMarkDown(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteMarkDown(blog *model.Blog) error {
	return WriteProcessedMarkdown(blog.FilePath, blog.BlogHeader, blog.BlogContent)
}
