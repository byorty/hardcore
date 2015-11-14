package criteria

import (
	"github.com/byorty/hardcore/db"
	"github.com/byorty/hardcore/expr"
	"github.com/byorty/hardcore/types"
)

type SelectCriteriaImpl struct {
	dao         types.DAO
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

func SelectByDAO(dao types.DAO) types.SelectCriteria {
	selectCriteria := newSelect()
	selectCriteria.dao = dao
	selectCriteria.proto = dao.Proto()
	return selectCriteria
}

func (s *SelectCriteriaImpl) And(logic types.Logic) types.SelectCriteria {
	return s.addLogic(types.AndLogicChainKind, logic)
}

func (s *SelectCriteriaImpl) addLogic(kind types.LogicChainKind, logic types.Logic) types.SelectCriteria {
	chain := expr.NewLogicChain(kind)
	chain.Add(logic)
	s.chains = append(s.chains, chain)

	if logic.GetArg() != nil {
		s.args = append(s.args, logic.GetArg())
	}
	return s
}

func (s *SelectCriteriaImpl) Or(logic types.Logic) types.SelectCriteria {
	return s.addLogic(types.OrLogicChainKind, logic)
}

func (s *SelectCriteriaImpl) One(model types.StraightMappingModel) {
	s.dao = model.DAO()
	s.proto = model.Proto()
	s.dao.One(s, model)
}

func (s *SelectCriteriaImpl) All(models types.StraightMappingModel) {
	s.dao = models.DAO()
	s.proto = models.Proto()
	s.dao.All(s, models)
}

func (s *SelectCriteriaImpl) ToNative() interface{} {
	writer := db.Pool().ByDAO(s.dao).GetQueryWriter()
	writer.SetProto(s.proto)
	writer.SetTable(s.dao.GetTable())
	writer.SetLogicChain(s.chains)
	writer.SetProjections(s.projections)
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
