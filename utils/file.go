package utils

import (
	"os"
	"github.com/byorty/hardcore/is"
)

func Pwd() string {
	pwd, err := os.Getwd()
	if is.Nil(err) {
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
