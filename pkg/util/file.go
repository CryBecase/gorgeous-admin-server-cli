package util

import (
	"os"
	"path/filepath"
)

// MakeNotExistDir 创建目录，如果此目录存在则直接返回
func MakeNotExistDir(path string, perm ...os.FileMode) error {
	if IsExist(path) {
		return nil
	}

	p := os.FileMode(0755)
	if len(perm) > 0 {
		p = perm[0]
	}

	err := os.MkdirAll(path, p)
	if err != nil {
		return err
	}

	return nil
}

// MakeNotExistFile 创建文件，如果此文件存在则直接返回
func MakeNotExistFile(path string) (*os.File, error) {
	if err := MakeNotExistDir(filepath.Dir(path)); err != nil {
		return nil, err
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.FileMode(0755))
	if err != nil {
		return nil, err
	}

	return f, nil
}

// IsExist 判断目录/文件是否存在
func IsExist(f string) bool {
	_, err := os.Stat(f)
	return err == nil || os.IsExist(err)
}
