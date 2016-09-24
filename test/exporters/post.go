package exporters

import (
	"github.com/byorty/hardcore/exporter"
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

type PostPropertyImpl struct {
	exporter.PropertyImpl
	closure func(*models.Post) interface{}
}

func newPostProperty(name string, closure func(post *models.Post) interface{}) types.ExportableProperty {
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

func NewPosts(posts models.Posts) types.Exporter {
	exp := new(exporter.BaseImpl)
	exp.SetProperties(postProperties)
	exp.SetExportable(posts)
	return exp
}

var (
	postProperties = []types.ExportableProperty{
		newPostProperty("id", func(post *models.Post) interface{} {
			return post.GetId()
		}),
		newPostProperty("name", func(post *models.Post) interface{} {
			return post.GetName()
		}),
		newPostProperty("description", func(post *models.Post) interface{} {
			return post.GetDescription()
		}),
	}
)
