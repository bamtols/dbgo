package cliFn

import (
	"os"
	"path/filepath"
	"strings"
)

func GetAbsolutePath(p string) string {
	if strings.HasPrefix(p, "/") {
		return p
	}
	pwd, _ := os.Getwd()
	return filepath.Join(pwd, p)
}

func GetAbsolutePathWithDefault(p *string, defPath string) string {
	pwd, _ := os.Getwd()
	if p == nil {
		return filepath.Join(pwd, defPath)
	} else {
		return GetAbsolutePath(*p)
	}
}

func CreateFilePath(fp string) error {
	if _, err := os.Stat(fp); os.IsExist(err) {
		return nil
	}
	return os.MkdirAll(fp, os.ModePerm)
}

func CreateFilePathByArray(list ...string) error {
	for _, fp := range list {
		if err := CreateFilePath(fp); err != nil {
			return err
		}
	}
	return nil
}
