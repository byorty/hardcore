package proj

type BaseImpl struct {
	name string
	table string
}

func (b *BaseImpl) IsWriteSqlPart() bool {
	return true
}
