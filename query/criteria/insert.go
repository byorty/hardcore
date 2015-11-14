package criteria

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/db"
)

type InserCriteriaImpl struct {
	dao         types.DAO
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

func (i *InserCriteriaImpl) One(model types.StraightMappingModel) {
	i.dao = model.DAO()
	i.proto = model.Proto()

	for _, property := range i.proto.GetSlice() {
		getter := property.GetGetter()
		i.AddArg(getter.Call(model))
	}

//	i.dao.One(i, model)
}

func (i *InserCriteriaImpl) All(models types.StraightMappingModel) {
	i.dao = models.DAO()
	i.proto = models.Proto()
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
