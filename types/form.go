package types

type Primitive interface {
	GetName() string
	GetError() string
	Required()
	Missing(string)
	Wrong(string)
	Custom(int, string)
	Import(interface{}) bool
	Export(interface{})
}

type Form interface {
	Add(Primitive) Form
	Check(RequestScope) (bool, map[string]string)
}
