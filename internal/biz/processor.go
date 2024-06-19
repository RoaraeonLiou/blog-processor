package biz

import (
	"blog-processor/global"
	"blog-processor/internal/db"
	"blog-processor/internal/model"
	"blog-processor/pkg/utils"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ProcessSingleFile(filePath string, httpBasePath string, imgOutPath string, dateLayout string) error {
	header, fileContent, err := utils.ReadYAMLHeader(filePath)
	if err != nil {
		return err
	}

	// 创建文件夹
	fileName := filepath.Base(filePath)
	dirPath := filepath.Dir(filePath)
	folderName := filepath.Base(dirPath)
	imageFolderName := utils.EncodeMD5(fileName)
	fatherFolderName := utils.EncodeMD5(folderName)
	newImageDirPath := imgOutPath + "/" + fatherFolderName + "/" + imageFolderName
	if err := os.MkdirAll(newImageDirPath, os.ModePerm); err != nil {
		return err
	}

	// todo: 处理头部
	utils.ProcessHeader(header, fileContent, fileName, dateLayout)

	imagesPaths := utils.ExtractImagesAndEncodeFilename(fileContent)

	newContent := utils.ReplaceImagePaths(fileContent, imagesPaths, httpBasePath, newImageDirPath)

	err = utils.CopyImagesToDir(imagesPaths, newImageDirPath, dirPath)
	if err != nil {
		return err
	}
	err = utils.WriteProcessedMarkdown(filePath, header, newContent)
	if err != nil {
		return err
	}
	return nil
}

func ProcessAll(filePath string, httpBasePath string, imgOutPath string, dateLayout string) error {
	mdFilePaths, err := utils.FindAllMarkdownFiles(filePath)
	if err != nil {
		return err
	}
	for _, filePath := range mdFilePaths {
		err = ProcessSingleFile(filePath, httpBasePath, imgOutPath, dateLayout)
		if err != nil {
			return err
		}
	}
	return nil
}

func Process(filePath string, httpBasePath string, imgOutPath string, dateLayout string) error {
	mdFilePaths, err := utils.FindAllMarkdownFiles(filePath)
	if err != nil {
		return err
	}
	var blogsHash []string
	now := time.Now()
	unifyModTime := now.Format(dateLayout)
	for _, filePath := range mdFilePaths {
		var blog *model.Blog
		blog, err = parseAndHandleImageSource(filePath, httpBasePath, imgOutPath, dateLayout)
		if err != nil {
			return err
		}

		err = blogHandler(blog, unifyModTime)
		if err != nil {
			return err
		}
		blogsHash = append(blogsHash, blog.Md5Path)
	}
	// 删除不存在的数据
	allHash, _, err := db.QueryAllMd5()
	if err != nil {
		return err
	}
	diff := utils.Difference(allHash, blogsHash)
	for _, hash := range diff {
		err = db.DeleteByMd5(hash)
		if err != nil {
			return err
		}
	}
	return nil
}

func parseAndHandleImageSource(filePath string, httpBasePath string, imgOutPath string, dateLayout string) (*model.Blog, error) {
	// 创建文件夹
	fileName := filepath.Base(filePath)
	dirPath := filepath.Dir(filePath)
	folderName := filepath.Base(dirPath)
	imageFolderName := utils.EncodeMD5(fileName)
	fatherFolderName := utils.EncodeMD5(folderName)
	newImageDirPath := imgOutPath + "/" + fatherFolderName + "/" + imageFolderName
	if err := os.MkdirAll(newImageDirPath, os.ModePerm); err != nil {
		return nil, err
	}

	// 读取文件内容替换图片源
	rawByteContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	rawContent := string(rawByteContent)
	imagesPaths := utils.ExtractImagesAndEncodeFilename(rawContent)
	mappingDir := filepath.Base(global.BasicSetting.OutputDir) + "/" + fatherFolderName + "/" + imageFolderName
	newContent := utils.ReplaceImagePaths(rawContent, imagesPaths, httpBasePath, mappingDir)

	err = utils.CopyImagesToDir(imagesPaths, newImageDirPath, dirPath)
	if err != nil {
		return nil, err
	}

	// 分离头部和主体
	header, fileContent, err := utils.SplitHeaderAndContent(newContent)
	if err != nil {
		return nil, err
	}

	// 单独处理头部
	header = utils.HeaderRebuild(header, fileContent, fileName, dateLayout)

	// 组装Blog对象
	blog := &model.Blog{}
	//var blog *model.Blog
	blog.BlogContent = fileContent
	blog.DirName = folderName
	blog.FileName = fileName
	blog.FilePath = filePath
	blog.Md5Path = utils.EncodeMD5(blog.DirName) + "/" + utils.EncodeMD5(blog.FileName)
	blog.BlogHeader = header
	checkInfo := blog.BlogHeader.Title +
		strings.Join(blog.BlogHeader.Categories, ",") +
		strings.Join(blog.BlogHeader.Tags, ",") +
		blog.BlogHeader.Status +
		blog.DirName +
		blog.FileName +
		blog.FilePath
	blog.Hash = utils.EncodeSha256(checkInfo)

	return blog, nil
}

func blogHandler(blog *model.Blog, modTime string) error {
	// 查询数据库是否存在
	exist, err := db.Exist(blog)
	if err != nil {
		return err
	}
	if !exist {
		// 新文件
		blog.BlogHeader.LastMod = modTime
		err = utils.WriteMarkDown(blog)
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
		err = utils.WriteMarkDown(blog)
		if err != nil {
			return err
		}
		return nil
	} else {
		// 维护创建时间
		blog.BlogHeader.Date = meta.CreatedAt
		blog.BlogHeader.LastMod = modTime
		err = utils.WriteMarkDown(blog)
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
