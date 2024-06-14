package main

import (
	"blog-processor/global"
	"blog-processor/internal/biz"
	"blog-processor/pkg/setting"
	"fmt"
	"os"
	"path/filepath"
)

func init() {
	err := setupSetting()
	if err != nil {
		fmt.Println("setup setting err:", err)
	}
}

func main() {
	fmt.Println("Start processing")
	err := biz.ProcessAll(global.BasicSetting.BlogDir, global.BasicSetting.HttpBasePath, global.BasicSetting.OutputDir)
	if err != nil {
		fmt.Println("process err:", err)
	}
	fmt.Println("Done!")
}

func setupSetting() error {
	// 创建设置对象
	appSetting, err := setting.NewSetting()
	if err != nil {
		return err
	}
	// 读取配置
	err = appSetting.ReadSection("Basic", &global.BasicSetting)
	if err != nil {
		return err
	}
	err = appSetting.ReadSection("LogStrategy", &global.LogStrategySetting)
	if err != nil {
		return err
	}
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	global.BasicSetting.ImageDir = filepath.Join(wd, global.BasicSetting.ImageDir)
	global.BasicSetting.OutputDir = filepath.Join(wd, global.BasicSetting.OutputDir)
	global.BasicSetting.BlogDir = filepath.Join(wd, global.BasicSetting.BlogDir)
	global.BasicSetting.TemplateFile = filepath.Join(wd, global.BasicSetting.TemplateFile)
	// 配置成功，返回空错误对象
	return nil
}
