package util

import (
	"errors"
	"os"
	"path/filepath"
)

// PathExists 判断路径是否存在 by os.Stat 工作目录下的路径
func PathExists(path string) (bool, error) {
	workDir, _ := os.Getwd()
	path = filepath.Join(workDir, path)
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}
		return false, errors.New("存在同名文件")
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
