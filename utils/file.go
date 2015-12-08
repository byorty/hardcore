package utils

import (
	"os"
)

func Pwd() string {
	pwd, err := os.Getwd()
	if err == nil {
		return pwd
	}
	return ""
}

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
