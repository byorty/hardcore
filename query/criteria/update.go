package criteria

import (
	"github.com/byorty/hardcore/pool"
	"github.com/byorty/hardcore/query/expr"
	"github.com/byorty/hardcore/query/proj"
	"github.com/byorty/hardcore/types"
)

type UpdateCriteriaImpl struct {
	dao         types.ModelDAO
	proto       types.Proto
	args        []interface{}
	chains      []types.LogicChain
	projections []types.Projection
}

func Update() types.UpdateCriteria {
	return newUpdate()
}

func newUpdate() *UpdateCriteriaImpl {
	return &UpdateCriteriaImpl{
		chains:      make([]types.LogicChain, 0),
		projections: make([]types.Projection, 0),
		args:        make([]interface{}, 0),
	}
}

func UpdateByDAO(dao types.ModelDAO) types.UpdateCriteria {
	update := newUpdate()

	update.dao = dao
	update.proto = dao.Proto()

	for _, property := range update.proto.GetSlice() {
		if property.GetRelation() == types.ProtoNoneRelation && property.GetField() != "id" {
			update.Add(proj.Eq(property.GetName(), nil))
		}
	}

	return update
}

func (u *UpdateCriteriaImpl) AddArg(arg interface{}) {
	u.args = append(u.args, arg)
}

func (u *UpdateCriteriaImpl) One(model types.Model) {
	u.dao = model.CommonDAO()
	u.proto = model.Proto()

	properties := u.proto.GetMap()
	for name, property := range properties {
		if property.GetRelation() == types.ProtoNoneRelation {
			getter := property.GetGetter()
			u.Add(proj.Eq(name, getter(model)))
		}
	}

	u.dao.Update(u, model)
}

func (u *UpdateCriteriaImpl) All(models types.Model) {
	//	i.dao = models.DAO()
	//	i.proto = models.Proto()
	//	i.dao.All(i, models)
}

func (u *UpdateCriteriaImpl) GetArgs() []interface{} {
	return u.args
}

func (u *UpdateCriteriaImpl) ToNative() interface{} {
	writer := pool.DB().ByDAO(u.dao).GetQueryWriter()
	writer.SetProto(u.proto)
	writer.SetTable(u.dao.GetTable())
	writer.SetLogicChain(u.chains)
	writer.SetProjections(u.projections)
	writer.SetArgs(u.args)

	query := writer.WriteUpdate()
	u.args = writer.GetArgs()
	return query
}

func (u *UpdateCriteriaImpl) And(logic types.Logic) types.UpdateCriteria {
	return u.addLogic(types.AndLogicChainKind, logic)
}

func (u *UpdateCriteriaImpl) addLogic(kind types.LogicChainKind, logic types.Logic) types.UpdateCriteria {
	logic.AddArg(u)
	chain := expr.NewLogicChain(kind)
	chain.Add(logic)
	u.chains = append(u.chains, chain)
	return u
}

func (u *UpdateCriteriaImpl) Or(logic types.Logic) types.UpdateCriteria {
	return u.addLogic(types.OrLogicChainKind, logic)
}

func (u *UpdateCriteriaImpl) Add(projection types.Projection) types.UpdateCriteria {
	u.projections = append(u.projections, projection)
	return u
}
