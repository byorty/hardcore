package exporters

import (
	"github.com/byorty/hardcore/types"
	"github.com/byorty/hardcore/exporter"
	"github.com/byorty/hardcore/test/models"
)

type PostPropertyImpl struct {
	exporter.PropertyImpl
	closure func(*models.Post) interface{}
}

func NewPostProperty(name string, closure func(post *models.Post) interface{}) types.ExportableProperty {
	return &PostPropertyImpl{
		exporter.NewProperty(name),
		closure,
	}
}

func (p PostPropertyImpl) GetValue(model interface{}) interface{} {
	return p.closure(model.(*models.Post))
}

func NewPost(post *models.Post) types.Exporter {
	exp := new(exporter.BaseImpl)
	exp.SetProperties(postProperties)
	exp.SetExportable(post)
	return exp
}

var (
	postProperties = []types.ExportableProperty{ 
		NewPostProperty("id", func(post *models.Post) interface{} {
			return post.GetId()
		}),
		NewPostProperty("name", func(post *models.Post) interface{} {
			return post.GetName()
		}),
		NewPostProperty("description", func(post *models.Post) interface{} {
			return post.GetDescription()
		}),
	}
)
