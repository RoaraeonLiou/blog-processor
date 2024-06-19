package main

import (
	"blog-processor/global"
	"blog-processor/internal/biz"
	"blog-processor/internal/db"
	"blog-processor/pkg/setting"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

func init() {
	err := setupSetting()
	if err != nil {
		fmt.Println("setup setting err:", err)
	}
	err = setupDBEngine()
	if err != nil {
		fmt.Println("setup db engine err:", err)
	}
}

func main() {
	defer func(LiteDB *sql.DB) {
		if LiteDB != nil {
			err := LiteDB.Close()
			if err != nil {
				fmt.Println("close db err:", err)
			}
		}
	}(global.LiteDB)
	fmt.Println("Start processing")
	//err := biz.ProcessAll(global.BasicSetting.BlogDir,
	//	global.BasicSetting.HttpBasePath,
	//	global.BasicSetting.OutputDir,
	//	global.BasicSetting.DateLayout)
	err := biz.Process(global.BasicSetting.BlogDir,
		global.BasicSetting.HttpBasePath,
		global.BasicSetting.OutputDir,
		global.BasicSetting.DateLayout)
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
	err = appSetting.ReadSection("DataBase", &global.DBSetting)
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

func setupDBEngine() error {
	var err error
	global.LiteDB, err = db.NewDBEngine(global.DBSetting)
	if err != nil {
		return err
	}
	global.Insert, err = global.LiteDB.Prepare(db.INSERT_META)
	if err != nil {
		return err
	}
	global.Update, err = global.LiteDB.Prepare(db.UPDATE_META)
	if err != nil {
		return err
	}
	global.QueryMeta, err = global.LiteDB.Prepare(db.QUERY_META)
	if err != nil {
		return err
	}
	global.QueryAll, err = global.LiteDB.Prepare(db.QUERY_META_ALL)
	if err != nil {
		return err
	}
	global.Delete, err = global.LiteDB.Prepare(db.DELETE_META)
	if err != nil {
		return err
	}
	global.Exists, err = global.LiteDB.Prepare(db.QUERY_EXIST)
	if err != nil {
		return err
	}
	return nil
}
