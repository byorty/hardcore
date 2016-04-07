package writer

import "github.com/byorty/hardcore/types"

type BaseImpl struct {
	proto       types.Proto
	table       string
	chains      []types.LogicChain
	projections []types.Projection
	orders      []types.Order
	args        []interface{}
	limit       int
	offset      int
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

func (b *BaseImpl) SetOrders(orders []types.Order) {
	b.orders = orders
}

func (b *BaseImpl) SetArgs(args []interface{}) {
	b.args = args
}

func (b *BaseImpl) AddArg(arg interface{}) {
	b.args = append(b.args, arg)
}

func (b *BaseImpl) GetArgs() []interface{} {
	return b.args
}

func (b *BaseImpl) SetLimit(limit int) {
	b.limit = limit
}

func (b *BaseImpl) SetOffset(offset int) {
	b.offset = offset
}
