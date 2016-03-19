package utils

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5(str string) string {
	h := md5.New()
	io.WriteString(h, string)
	return fmt.Sprintf("%x", h.Sum(nil))
}
