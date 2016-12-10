package types

type ModelScanner interface {
	One(Model)
	All(Model)
}

type EnumScanner interface {
	Eq(int) bool
	Scan(int)
	One(Named)
}
