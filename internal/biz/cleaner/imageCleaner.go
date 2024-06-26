package cleaner

import (
	"blog-processor/global"
	"fmt"
	"os"
	"path/filepath"
)

func DeleteAllImages() error {
	imageExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
		".tiff": true,
		".webp": true,
	}

	err := filepath.Walk(global.BasicSetting.BlogDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 检查是否为文件，并且扩展名是否为图片格式
		if !info.IsDir() && imageExtensions[filepath.Ext(path)] {
			if err := os.Remove(path); err != nil {
				return err
			}
			fmt.Printf("Deleted image file: %s\n", path)
		}
		return nil
	})

	return err
}
