package util

import "os"

func DoesFileExist(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}
