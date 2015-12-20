package utils

import (
    "strings"
)


func LowerFirst(str string) string {
    if str == "" {
        return str
    }
    return strings.ToLower(str[0:1]) + str[1:]
}

func UpperFirst(str string) string {
    if str == "" {
        return str
    }
    return strings.ToUpper(str[0:1]) + str[1:]
}