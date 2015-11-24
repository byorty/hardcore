package proj

type BaseImpl struct {
	name string
	alias string
}

func (b *BaseImpl) Alias(alias string) {
	b.alias = alias
}
