package model

type Blog struct {
	Md5Path     string      `json:"md5_path"`
	BlogHeader  *BlogHeader `json:"blog_header"`
	BlogContent string      `json:"blog_content"`
	FileName    string      `json:"file_name"`
	DirName     string      `json:"dir_name"`
	FilePath    string      `json:"file_path"`
	Hash        string      `json:"hash"`
}
