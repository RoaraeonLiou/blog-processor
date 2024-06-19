package db

var (
	//CREATE_TABLE = `CREATE TABLE IF NOT EXISTS blog_meta(
	//	md5_path VARCHAR(256) PRIMARY KEY,
	//	title VARCHAR(256) NOT NULL,
	//	created_at VARCHAR(64) NOT NULL,
	//	updated_at VARCHAR(64) NOT NULL,
	//	status VARCHAR(64) NOT NULL
	//);`
	CREATE_TABLE = `CREATE TABLE IF NOT EXISTS blog_meta(	
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    md5_path VARCHAR(256) NOT NULL,
    created_at VARCHAR(64) NOT NULL,
    updated_at VARCHAR(64) NOT NULL,
    hash VARCHAR(300) NOT NULL,
    CONSTRAINT idx_md5_path UNIQUE(md5_path)
	);`
	INSERT_META    = `INSERT INTO blog_meta(md5_path, created_at, updated_at, hash) VALUES (?, ?, ?, ?)`
	UPDATE_META    = `UPDATE blog_meta SET hash = ?, updated_at = ? WHERE md5_path = ?`
	DELETE_META    = `DELETE FROM blog_meta WHERE md5_path = ?`
	QUERY_META_ALL = `SELECT * FROM blog_meta WHERE md5_path = ?`
	QUERY_META     = `SELECT created_at, updated_at, hash FROM blog_meta WHERE md5_path = ?`
	QUERY_EXIST    = `SELECT EXISTS(SELECT 1 FROM blog_meta WHERE md5_path = ? LIMIT 1)`
	QUERY_ALL_MD5  = `SELECT md5_path FROM blog_meta`
)
