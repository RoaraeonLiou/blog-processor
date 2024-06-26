package biz

import (
	"blog-processor/global"
	"blog-processor/internal/biz/cleaner"
	"blog-processor/internal/biz/processor"
	"blog-processor/internal/db"
	"blog-processor/pkg/utils/set_operation"
	"fmt"
)

// 总体调度

func Exec() {
	// 迭代处理各个子文件夹
	//folders, err := ScanFolder(global.BasicSetting.BlogDir)
	//if err != nil {
	//	fmt.Println("处理文件夹出错")
	//	return
	//}
	folders, commonHeadMap, err := ScanFolderWithCommonHeader(global.BasicSetting.BlogDir)
	if err != nil {
		fmt.Println("处理文件夹出错")
		return
	}
	var md5Set []string
	for _, folder := range folders {
		subfolderMd5Set, err := processor.ProcessFolder(folder, commonHeadMap[folder])
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
	diff := set_operation.Difference(oldMd5Set, md5Set)
	for _, hash := range diff {
		err = db.DeleteByMd5(hash)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// 删除原有图片文件
	err = cleaner.DeleteAllImages()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 删除空目录
	for _, folder := range folders {
		err := cleaner.DeleteEmptyDir(folder)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// 添加模版文件
	AddArchivesFile()
	AddSearchFile()
}
