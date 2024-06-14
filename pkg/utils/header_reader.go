package utils

import (
	"blog-processor/internal/model"
	"bufio"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

func ReadYAMLHeader(filePath string) (*model.BlogHeader, string, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var yamlLines, bodyContent string
	readingYAML, readYAML := false, false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "---" && !readYAML {
			// 如果正在读取YAML, 此时第二次遇到“---”表示结束, 设置readYAML为true
			if readingYAML {
				readYAML = true
				readingYAML = false
				continue
			}
			// 如果还没有读取完YAML, 此时第一次遇到“---”表示开始读取YAML, 设置readingTAML为true
			if !readYAML {
				readingYAML = true
				continue
			}
		}
		if readingYAML {
			yamlLines += line + "\n"
		} else {
			bodyContent += line + "\n"
		}
	}
	var header model.BlogHeader
	err = yaml.Unmarshal([]byte(yamlLines), &header)
	if err != nil {
		return nil, "", "", err
	}

	fileName := filepath.Base(filePath)
	return &header, fileName, bodyContent, nil
}
