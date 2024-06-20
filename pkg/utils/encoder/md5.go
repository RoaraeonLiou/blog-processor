package encoder

import (
	"crypto/md5"
	"encoding/hex"
)

func EncodeMD5(value string) string {
	// 用于针对上传后的文件名进行格式化, 防止原始名称暴露
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
