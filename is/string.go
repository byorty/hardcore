package is

func EmptyStr(str string) bool {
	return str == ""
}

func NotEmptyStr(str string) bool {
	return !EmptyStr(str)
}
