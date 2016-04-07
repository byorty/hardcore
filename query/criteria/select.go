package criteria

import (
	"github.com/byorty/hardcore/pool"
	"github.com/byorty/hardcore/query/expr"
	"github.com/byorty/hardcore/types"
)

type SelectCriteriaImpl struct {
	dao         types.ModelDAO
	proto       types.Proto
	chains      []types.LogicChain
	projections []types.Projection
	orders      []types.Order
	args        []interface{}
	limit       int
	offset      int
}

func Select() types.SelectCriteria {
	return newSelect()
}

func newSelect() *SelectCriteriaImpl {
	return &SelectCriteriaImpl{
		chains:      make([]types.LogicChain, 0),
		projections: make([]types.Projection, 0),
		orders:      make([]types.Order, 0),
		args:        make([]interface{}, 0),
	}
}

func SelectByDAO(dao types.ModelDAO) types.SelectCriteria {
	selectCriteria := newSelect()
	selectCriteria.dao = dao
	selectCriteria.proto = dao.Proto()
	return selectCriteria
}

func (s *SelectCriteriaImpl) And(logic types.Logic) types.SelectCriteria {
	return s.addLogic(types.AndLogicChainKind, logic)
}

func (s *SelectCriteriaImpl) addLogic(kind types.LogicChainKind, logic types.Logic) types.SelectCriteria {
	logic.AddArg(s)
	chain := expr.NewLogicChain(kind)
	chain.Add(logic)
	s.chains = append(s.chains, chain)
	return s
}

func (s *SelectCriteriaImpl) AddArg(arg interface{}) {
	s.args = append(s.args, arg)
}

func (s *SelectCriteriaImpl) Or(logic types.Logic) types.SelectCriteria {
	return s.addLogic(types.OrLogicChainKind, logic)
}

func (s *SelectCriteriaImpl) One(model types.Model) {
	s.dao = model.CommonDAO()
	s.proto = model.Proto()
	s.dao.SelectOne(s, model)
}

func (s *SelectCriteriaImpl) All(models types.Model) {
	s.dao = models.CommonDAO()
	s.proto = models.Proto()
	s.dao.SelectAll(s, models)
}

func (s *SelectCriteriaImpl) ToNative() interface{} {
	writer := pool.DB().ByDAO(s.dao).GetQueryWriter()
	writer.SetProto(s.proto)
	writer.SetTable(s.dao.GetTable())
	writer.SetLogicChain(s.chains)
	writer.SetProjections(s.projections)
	writer.SetOrders(s.orders)
	writer.SetArgs(s.args)
	writer.SetLimit(s.limit)
	writer.SetOffset(s.offset)
	return writer.WriteSelect()
}

func (s *SelectCriteriaImpl) GetArgs() []interface{} {
	return s.args
}

func (s *SelectCriteriaImpl) Custom(items ...interface{}) {
	s.dao.Custom(s.dao, s, items...)
}

func (s *SelectCriteriaImpl) Add(projection types.Projection) types.SelectCriteria {
	s.projections = append(s.projections, projection)
	return s
}

func (s *SelectCriteriaImpl) Limit(limit int) types.SelectCriteria {
	s.limit = limit
	return s
}

func (s *SelectCriteriaImpl) Offset(offset int) types.SelectCriteria {
	s.offset = offset
	return s
}

func (s *SelectCriteriaImpl) Order(order types.Order) types.SelectCriteria {
	s.orders = append(s.orders, order)
	return s
}
