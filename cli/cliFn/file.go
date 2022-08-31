package cliFn

import (
	"os"
)

func CreateFile(fp string) (*os.File, error) {
	if _, err := os.Stat(fp); os.IsExist(err) {
		if err := os.Remove(fp); err != nil {
			return nil, err
		}
	}

	return os.Create(fp)
}
