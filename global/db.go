package global

import (
	"database/sql"
)

var (
	LiteDB    *sql.DB
	Insert    *sql.Stmt
	Update    *sql.Stmt
	Delete    *sql.Stmt
	QueryAll  *sql.Stmt
	Exists    *sql.Stmt
	QueryMeta *sql.Stmt
)
