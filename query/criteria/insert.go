package criteria

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/orm/db"
)

type InserCriteriaImpl struct {
	dao         types.ModelDAO
	proto       types.Proto
	args        []interface{}
}

func Insert() types.Criteria {
	return newInsert()
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
	i.dao = model.DAO()
	i.proto = model.Proto()

	for _, property := range i.proto.GetSlice() {
		if property.GetField() != "id" {
			getter := property.GetGetter()
			i.AddArg(getter.Call(model))
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
	writer := db.Pool().ByDAO(i.dao).GetQueryWriter()
	writer.SetProto(i.proto)
	writer.SetTable(i.dao.GetTable())
	writer.SetArgs(i.args)
	return writer.WriteInsert()
}
