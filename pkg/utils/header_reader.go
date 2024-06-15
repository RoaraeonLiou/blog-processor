package utils

import (
	"blog-processor/internal/model"
	"bufio"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"time"
)

func ReadYAMLHeader(filePath string) (*model.BlogHeader, string, error) {
	// 分离头部和主体内容, 后续分别做处理
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", err
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
		return nil, "", err
	}

	return &header, bodyContent, nil
}

func extractFirstHeading(content string) string {
	content = strings.TrimSpace(content)
	lines := strings.SplitN(content, "\n", 2)
	if len(lines) > 0 {
		firstLine := lines[0]
		if strings.HasPrefix(firstLine, "# ") {
			return strings.TrimSpace(firstLine[2:])
		}
	}
	return ""
}

func ProcessHeader(header *model.BlogHeader, bodyContent string, fileName string, dateLayout string) error {
	if header.Title == "" {
		header.Title = extractFirstHeading(bodyContent)
		if header.Title == "" {
			header.Title = fileName
		}
	}
	if header.Date == "" {
		now := time.Now()
		formattedDate := now.Format(dateLayout)
		header.Date = formattedDate
	}
	return nil
}
