package db

import (
	"blog-processor/pkg/setting"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func NewDBEngine(dbSetting *setting.DBSettingS) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbSetting.DBFile)
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(CREATE_TABLE)
	if err != nil {
		return nil, err
	}

	return db, nil
}
