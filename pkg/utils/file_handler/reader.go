package file_handler

import "os"

func ReadMarkDown(filePath string) (string, int64, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", -1, err
	}
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return "", -1, err
	}
	fileSize := fileInfo.Size()
	return string(data), fileSize, nil
}
