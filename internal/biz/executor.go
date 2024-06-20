package biz

import (
	"blog-processor/global"
	"blog-processor/internal/db"
	"blog-processor/pkg/utils"
	"fmt"
)

// 总体调度

func Exec() {
	// 迭代处理各个子文件夹
	folders, err := ScanFolder(global.BasicSetting.BlogDir)
	if err != nil {
		fmt.Println("处理文件夹出错")
		return
	}
	var md5Set []string
	for _, folder := range folders {
		subfolderMd5Set, err := ProcessFolder(folder)
		if err != nil {
			fmt.Println(err)
			return
		}
		md5Set = append(md5Set, subfolderMd5Set...)
	}

	// 删除不存在的数据
	oldMd5Set, _, err := db.QueryAllMd5()
	if err != nil {
		fmt.Println(err)
		return
	}
	diff := utils.Difference(oldMd5Set, md5Set)
	for _, hash := range diff {
		err = db.DeleteByMd5(hash)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// 删除原有图片文件
	err = DeleteAllImages()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 删除空目录
	for _, folder := range folders {
		err := DeleteEmptyDir(folder)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// 添加模版文件
	AddArchivesFile()
	AddSearchFile()
}
