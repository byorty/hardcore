package utils

import (
	"github.com/byorty/hardcore/is"
	"os"
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
