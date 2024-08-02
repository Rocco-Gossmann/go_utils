package go_utils

import "os"

func MkDir(dirPath string) error {

	info, err := os.Stat(dirPath)

	if os.IsNotExist(err) {
		return os.MkdirAll(dirPath, os.ModePerm)
	} else if !info.IsDir() {
		Panicf("'%s' exists and is not a Directorys", dirPath)
	}

	return nil
}
