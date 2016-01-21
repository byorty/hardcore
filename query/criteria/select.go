package criteria

import (
	"github.com/byorty/hardcore/query/expr"
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/pool"
)

type SelectCriteriaImpl struct {
	dao         types.ModelDAO
	proto       types.Proto
	chains      []types.LogicChain
	projections []types.Projection
	args        []interface{}
}

func Select() types.SelectCriteria {
	return newSelect()
}

func newSelect() *SelectCriteriaImpl {
	return &SelectCriteriaImpl{
		chains:      make([]types.LogicChain, 0),
		projections: make([]types.Projection, 0),
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
	s.dao.One(s, model)
}

func (s *SelectCriteriaImpl) All(models types.Model) {
	s.dao = models.CommonDAO()
	s.proto = models.Proto()
	s.dao.All(s, models)
}

func (s *SelectCriteriaImpl) ToNative() interface{} {
	writer := pool.DB().ByDAO(s.dao).GetQueryWriter()
	writer.SetProto(s.proto)
	writer.SetTable(s.dao.GetTable())
	writer.SetLogicChain(s.chains)
	writer.SetProjections(s.projections)
	writer.SetArgs(s.args)
	return writer.WriteSelect()
}

func (s *SelectCriteriaImpl) GetArgs() []interface{} {
	return s.args
}

func (s *SelectCriteriaImpl) Custom(items ...interface{}) {
	s.dao.Custom(s.dao, s, items...)
}

//func (s *SelectCriteriaImpl) Customs(items ...[]interface{}) {
//
//}

func (s *SelectCriteriaImpl) Add(projection types.Projection) types.SelectCriteria {
	s.projections = append(s.projections, projection)
	return s
}
