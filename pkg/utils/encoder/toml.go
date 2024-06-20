package encoder

import (
	"bytes"
	"github.com/BurntSushi/toml"
)

func EncodeToTOMLString(data interface{}) (string, error) {
	var buf bytes.Buffer
	encoder := toml.NewEncoder(&buf)
	if err := encoder.Encode(data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
