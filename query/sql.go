package query

type Sql []interface{}

func (s Sql) ToNative() interface{} {
	if len(s) > 1 {
		return s[0]
	} else {
		return nil
	}
}

func (s Sql) GetArgs() []interface{} {
	if len(s) >= 2 {
		return s[1:]
	} else {
		return nil
	}
}
