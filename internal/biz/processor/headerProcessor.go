package processor

import (
	"blog-processor/global"
	"blog-processor/internal/model"
	"encoding/json"
	"errors"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
	"strings"
	"time"
)

func SplitHeaderAndBody(content string) (*model.BlogHeader, string, error) {
	var header model.BlogHeader

	// 检查并解析JSON头部
	if strings.HasPrefix(content, "{") && strings.Contains(content, "}") {
		end := strings.Index(content, "}") + 1
		headerContent := content[:end]
		if err := json.Unmarshal([]byte(headerContent), &header); err != nil {
			return nil, "", err
		}
		bodyContent := strings.TrimSpace(content[end:])
		return &header, bodyContent, nil
	}

	// 检查并解析YAML头部
	if strings.HasPrefix(content, "---\n") && strings.Contains(content, "---\n") {
		start := strings.Index(content, "---\n") + 4
		end := strings.Index(content[start:], "---\n") + start
		headerContent := content[start:end]
		if err := yaml.Unmarshal([]byte(headerContent), &header); err != nil {
			return nil, "", err
		}
		bodyContent := strings.TrimSpace(content[end+4:])
		return &header, bodyContent, nil
	}

	// 检查并解析TOML头部
	if strings.HasPrefix(content, "+++\n") && strings.Contains(content, "+++\n") {
		start := strings.Index(content, "+++\n") + 4
		end := strings.Index(content[start:], "+++\n") + start
		headerContent := content[start:end]
		if _, err := toml.Decode(headerContent, &header); err != nil {
			return nil, "", err
		}
		bodyContent := strings.TrimSpace(content[end+4:])
		return &header, bodyContent, nil
	}

	// 如果没有头部，返回空头部和完整内容
	return new(model.BlogHeader), content, nil
}

func HeaderRebuild(header *model.BlogHeader, fileContent string, fileName string) error {
	if header.Title == "" {
		header.Title = extractFirstHeading(fileContent)
		if header.Title == "" {
			header.Title = fileName
		}
	}

	if header.Date == "" {
		now := time.Now()
		formattedDate := now.Format(global.BasicSetting.DateLayout)
		header.Date = formattedDate
	}
	return nil
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

func ParseHeader(content string, ext string, format string) (*model.BlogHeader, error) {
	var header model.BlogHeader
	if ext == ".md" {
		if format == "yaml" || format == "yml" {
			if strings.HasPrefix(content, "---\n") && strings.Contains(content, "---\n") {
				start := strings.Index(content, "---\n") + 4
				end := strings.Index(content[start:], "---\n") + start
				content = content[start:end]
			}
			if err := yaml.Unmarshal([]byte(content), &header); err != nil {
				return nil, err
			}
		} else if format == "json" {
			if strings.HasPrefix(content, "{") && strings.Contains(content, "}") {
				end := strings.Index(content, "}") + 1
				content = content[:end]
			}
			if err := json.Unmarshal([]byte(content), &header); err != nil {
				return nil, err
			}
		} else if format == "toml" {
			if strings.HasPrefix(content, "+++\n") && strings.Contains(content, "+++\n") {
				start := strings.Index(content, "+++\n") + 4
				end := strings.Index(content[start:], "+++\n") + start
				content = content[start:end]
			}
			if err := toml.Unmarshal([]byte(content), &header); err != nil {
				return nil, err
			}
		}
	} else if (ext == ".yaml" || ext == ".yml") && (format == "yaml" || format == "yml") {
		if err := yaml.Unmarshal([]byte(content), &header); err != nil {
			return nil, err
		}
	} else if (ext == ".json") && format == "json" {
		if err := json.Unmarshal([]byte(content), &header); err != nil {
			return nil, err
		}
	} else if (ext == ".toml") && format == "toml" {
		if err := toml.Unmarshal([]byte(content), &header); err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("config error")
	}
	return &header, nil
}
