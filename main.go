package main

import (
	"blog-processor/global"
	"blog-processor/pkg/setting"
	"fmt"
)

func init() {
	err := setupSetting()
	if err != nil {
		fmt.Println("setup setting err:", err)
	}
}

func main() {
	fmt.Println("Start processing")
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
	// 配置成功，返回空错误对象
	return nil
}
