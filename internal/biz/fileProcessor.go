package biz

import (
	"blog-processor/global"
	"blog-processor/internal/db"
	"blog-processor/internal/model"
	"blog-processor/pkg/utils/encoder"
	"blog-processor/pkg/utils/file_handler"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func ProcessFile(filePath string) (string, error) {
	// 读取文件并替换文件源
	rawContent, fileSize, err := file_handler.ReadMarkDown(filePath)
	if err != nil {
		return "", err
	}
	imageProcessor := ImageProcessor{
		Content:    rawContent,
		Filepath:   filePath,
		OutputPath: "",
		BlogImages: []model.BlogImage{},
	}
	err = imageProcessor.ProcessImage()
	if err != nil {
		return "", err
	}

	// 分离头部与主体
	sourceHeader, content, err := SplitHeaderAndBody(imageProcessor.Content)
	if err != nil {
		return "", err
	}

	// 处理头部
	fileFullName := filepath.Base(filePath)
	fileExt := filepath.Ext(fileFullName)
	fileName := strings.TrimSuffix(fileFullName, fileExt)
	err = HeaderRebuild(sourceHeader, content, fileName)
	if err != nil {
		return "", err
	}

	// 构建blog对象
	blog := &model.Blog{}
	blog.BlogContent = content
	blog.FilePath = filePath
	blog.RelPath, _ = filepath.Rel(global.BasicSetting.BlogDir, filePath)
	blog.FileName = filepath.Base(filePath)
	blog.Md5Path = encoder.EncodeMD5(blog.RelPath)
	blog.BlogHeader = sourceHeader
	checkInfo := blog.BlogHeader.Title +
		strings.Join(blog.BlogHeader.Categories, ",") +
		strings.Join(blog.BlogHeader.Tags, ",") +
		blog.BlogHeader.Status +
		blog.RelPath +
		strconv.FormatInt(fileSize, 10)
	blog.Hash = encoder.EncodeSha256(checkInfo)

	// 写入文件
	err = WriteFile(blog)
	if err != nil {
		return "", err
	}
	return blog.Md5Path, nil
}

func WriteFile(blog *model.Blog) error {
	// 查询数据库是否存在
	exist, err := db.Exist(blog)
	if err != nil {
		return err
	}
	now := time.Now()
	modTime := now.Format(global.BasicSetting.DateLayout)
	if !exist {
		// 新文件
		blog.BlogHeader.LastMod = modTime
		err = file_handler.WriteMarkDown(blog)
		if err != nil {
			return err
		}
		err = db.Insert(blog)
		if err != nil {
			return err
		}
		return nil
	}
	// 读取元数据
	meta, err := db.QueryMeta(blog)
	if err != nil {
		return err
	}
	if meta.Hash == blog.Hash {
		// 文件没有修改, 不操作数据库, 直接写文件
		err = file_handler.WriteMarkDown(blog)
		if err != nil {
			return err
		}
		return nil
	} else {
		// 维护创建时间
		blog.BlogHeader.Date = meta.CreatedAt
		blog.BlogHeader.LastMod = modTime
		err = file_handler.WriteMarkDown(blog)
		if err != nil {
			return err
		}
		err = db.Update(blog)
		if err != nil {
			return err
		}
		return nil
	}
}
