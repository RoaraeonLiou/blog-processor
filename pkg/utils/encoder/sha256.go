package encoder

import (
	"crypto/sha256"
	"fmt"
)

func EncodeSha256(data string) string {
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}
