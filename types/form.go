package types

type Primitive interface {
	Import(interface{})
	Export() interface{}
}

type Form interface {
	Add(Primitive) Form
}
