package fsFn

import "os"

func IsExistFile(fp string) bool {
	if _, err := os.Stat(fp); os.IsExist(err) {
		return true
	}
	return false
}

func RemoveFile(fp string) error {
	return os.Remove(fp)
}
