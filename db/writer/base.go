package writer

import "github.com/byorty/hardcore/types"

type BaseImpl struct {
	proto types.Proto
	table string
	conditions []types.Logic
}

func (b *BaseImpl) SetProto(proto types.Proto) {
	b.proto = proto
}

func (b *BaseImpl) SetTable(table string) {
	b.table = table
}

func (b *BaseImpl) SetConditions(conditions []types.Logic) {
	b.conditions = conditions
}
