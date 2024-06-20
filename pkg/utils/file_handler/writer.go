package file_handler

import (
	"blog-processor/internal/model"
	"blog-processor/pkg/utils/encoder"
	"bufio"
	"os"
)

func WriteProcessedMarkdown(filePath string, header *model.BlogHeader, bodyContent string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString("+++\n")
	if err != nil {
		return err
	}

	//headerString, err := yaml.Marshal(header)
	headerString, err := encoder.EncodeToTOMLString(header)
	if err != nil {
		return err
	}

	_, err = writer.Write([]byte(headerString))
	if err != nil {
		return err
	}
	_, err = writer.WriteString("+++\n")
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

func WriteMarkDown(blog *model.Blog) error {
	return WriteProcessedMarkdown(blog.FilePath, blog.BlogHeader, blog.BlogContent)
}
