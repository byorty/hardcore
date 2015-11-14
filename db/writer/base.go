package writer

import "github.com/byorty/hardcore/types"

type BaseImpl struct {
	proto types.Proto
	table string
	chains []types.LogicChain
	projections []types.Projection
}

func (b *BaseImpl) SetProto(proto types.Proto) {
	b.proto = proto
}

func (b *BaseImpl) SetTable(table string) {
	b.table = table
}

func (b *BaseImpl) SetLogicChain(chains []types.LogicChain) {
	b.chains = chains
}

func (b *BaseImpl) SetProjections(projections []types.Projection) {
	b.projections = projections
}
