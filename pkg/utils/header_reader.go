package utils

import (
	"blog-processor/internal/model"
	"bufio"
	"errors"
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

func SplitHeaderAndContent(stringContent string) (*model.BlogHeader, string, error) {
	parts := strings.SplitN(stringContent, "---\n", 3)
	var yamlString string
	var bodyContent string
	var header model.BlogHeader
	switch len(parts) {
	case 1:
		// 没有“---”分割
		yamlString = parts[0]
		err := yaml.Unmarshal([]byte(yamlString), &header)
		if err != nil {
			// 文件没有头部
			bodyContent = parts[0]
			return nil, bodyContent, nil
		} else {
			// 文件只有头部
			return &header, "", nil
		}
	case 2:
		// 只有一个“---”分割
		if strings.TrimSpace(parts[0]) != "" {
			yamlString = parts[0]
			bodyContent = parts[1]
			err := yaml.Unmarshal([]byte(yamlString), &header)
			if err != nil {
				// “---”是主体内容一部分, 文件没有头部
				bodyContent = parts[0] + "---\n" + parts[1]
				return nil, bodyContent, nil
			} else {
				// "---"是不规范的头部分割, 但是可以分离出头部和主体
				return &header, bodyContent, nil
			}
		} else {
			yamlString = parts[1]
			err := yaml.Unmarshal([]byte(yamlString), &header)
			if err != nil {
				bodyContent = parts[0] + "---\n" + parts[1]
				return nil, bodyContent, nil
			} else {
				// “---”是不规范的头部分割, 文件只有头部
				return &header, "", nil
			}
		}
	case 3:
		// 至少有两个“---”分割
		if strings.TrimSpace(parts[0]) != "" {
			// 第一个分割出来的部分不是空串
			yamlString = parts[0]
			err := yaml.Unmarshal([]byte(yamlString), &header)
			if err != nil {
				// 认为没有头部
				return nil, stringContent, err
			}
		} else {
			yamlString = parts[1]
			bodyContent = parts[2]
			err := yaml.Unmarshal([]byte(yamlString), &header)
			if err != nil {
				// 认为是文件主体内容
				return nil, stringContent, err
			}
			return &header, bodyContent, nil
		}
	}

	return nil, "", errors.New("parse error")
}

func HeaderRebuild(originHeader *model.BlogHeader, bodyContent string, fileName string, dateLayout string) *model.BlogHeader {
	if originHeader == nil {
		originHeader = new(model.BlogHeader)
	}
	// 标题部分
	if originHeader.Title == "" {
		originHeader.Title = extractFirstHeading(bodyContent)
		if originHeader.Title == "" {
			originHeader.Title = fileName
		}
	}
	// 创建时间和修改时间部分
	if originHeader.Date == "" {
		now := time.Now()
		formattedDate := now.Format(dateLayout)
		originHeader.Date = formattedDate
	}
	if originHeader.LastMod == "" {
		now := time.Now()
		formattedDate := now.Format(dateLayout)
		originHeader.LastMod = formattedDate
	}
	// 状态部分
	if originHeader.Status == "" {
		originHeader.Status = "New"
	}
	return originHeader
}
