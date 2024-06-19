package model

type BlogMeta struct {
	Id        int    `json:"id"`
	MD5Path   string `json:"md5_path"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Hash      string `json:"hash"`
}
