package importers

import (
	"github.com/byorty/hardcore/importer"
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

type PostPropertyImpl struct {
	importer.PropertyImpl
	closure func(*models.Post, interface{})
}

func newPostProperty(kind types.ProtoKind, closure func(post *models.Post, value interface{})) types.ImportableProperty {
	return &PostPropertyImpl{
		importer.NewProperty(types.ScalarImportablePropertyKind, kind),
		closure,
	}
}

func (p PostPropertyImpl) SetValue(model interface{}, value interface{}) {
	p.closure(model.(*models.Post), value)
}

func NewPost(post *models.Post) types.Importer {
	imp := new(importer.BaseImpl)
	imp.SetProperties(postProperties)
	imp.SetImportable(post)
	return imp
}

func NewPosts(posts models.Posts) types.Importer {
	imp := new(importer.BaseImpl)
	imp.SetProperties(postProperties)
	imp.SetImportable(posts)
	return imp
}

var (
	postProperties = types.ImportableProperties{
		"id": newPostProperty(types.ProtoInt64Kind, func(post *models.Post, value interface{}) {
			post.SetId(value.(int64))
		}),
		"name": newPostProperty(types.ProtoStringKind, func(post *models.Post, value interface{}) {
			post.SetName(value.(string))
		}),
		"description": newPostProperty(types.ProtoStringKind, func(post *models.Post, value interface{}) {
			post.SetDescription(value.(string))
		}),
	}
)
