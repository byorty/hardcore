package is

func Nil(val interface{}) bool {
	return val == nil
}

func NotNil(val interface{}) bool {
	return !Nil(val)
}

func Eq(a, b interface{}) bool {
	return a == b
}

func NotEq(a, b interface{}) bool {
	return !Eq(a, b)
}
