package criteria

import (
	"github.com/byorty/hardcore/pool"
	"github.com/byorty/hardcore/types"
)

type InserCriteriaImpl struct {
	dao   types.ModelDAO
	proto types.Proto
	args  []interface{}
}

func Insert() types.Criteria {
	return newInsert()
}

func InsertByDao(dao types.ModelDAO) types.Criteria {
	insert := newInsert()
	insert.dao = dao
	insert.proto = dao.Proto()
	for _, property := range insert.proto.GetSlice() {
		if property.GetField() != "id" && property.GetRelation() == types.ProtoNoneRelation {
			insert.AddArg(nil)
		}
	}
	return insert
}

func newInsert() *InserCriteriaImpl {
	return &InserCriteriaImpl{
		args: make([]interface{}, 0),
	}
}

func (i *InserCriteriaImpl) AddArg(arg interface{}) {
	i.args = append(i.args, arg)
}

func (i *InserCriteriaImpl) One(model types.Model) {
	i.dao = model.CommonDAO()
	i.proto = model.Proto()

	for _, property := range i.proto.GetSlice() {
		if property.GetField() != "id" && property.GetRelation() == types.ProtoNoneRelation {
			getter := property.GetGetter()
			i.AddArg(getter(model))
		}
	}

	i.dao.Insert(i, model)
}

func (i *InserCriteriaImpl) All(models types.Model) {
	//	i.dao = models.DAO()
	//	i.proto = models.Proto()
	//	i.dao.All(i, models)
}

func (i *InserCriteriaImpl) GetArgs() []interface{} {
	return i.args
}

func (i *InserCriteriaImpl) ToNative() interface{} {
	writer := pool.DB().ByDAO(i.dao).GetQueryWriter()
	writer.SetProto(i.proto)
	writer.SetTable(i.dao.GetTable())
	writer.SetArgs(i.args)
	return writer.WriteInsert()
}
