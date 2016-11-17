package exporters

import (
	"github.com/byorty/hardcore/exporter"
	"github.com/byorty/hardcore/is"
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

type _PostImpl struct {
	item  *models.Post
	items models.Posts
	props []_PostPropertyImpl
	kind  types.ProtoKind
}

func NewPost(model *models.Post) types.Exporter {
	return _newPost(model, postProperties)
}

func _newPost(model *models.Post, props []_PostPropertyImpl) types.Exporter {
	exp := new(_PostImpl)
	exp.props = props
	exp.item = model
	exp.kind = types.ProtoModelKind
	return exp
}

func (p _PostImpl) Len() int {
	if is.Eq(p.kind, types.ProtoModelKind) {
		return len(p.props)
	} else {
		return p.items.Len()
	}
}

func (p _PostImpl) Get(i int) types.ExportableProperty {
	return p.props[i]
}

func (p _PostImpl) GetProtoKind() types.ProtoKind {
	return p.kind
}

func (p _PostImpl) Export(i int, encoder types.Encoder) {
	if is.Eq(p.kind, types.ProtoModelKind) {
		p.props[i].closure(p.item, encoder)
	} else {
		encoder.Encode(_newPost(p.items.Get(i), postProperties))
	}
}

func NewPosts(posts models.Posts) types.Exporter {
	exp := new(_PostImpl)
	exp.props = postProperties
	exp.items = posts
	exp.kind = types.ProtoModelSliceKind
	return exp
}

type _PostPropertyImpl struct {
	exporter.PropertyImpl
	closure func(*models.Post, types.Encoder)
}

func newPostProperty(name string, kind types.ProtoKind, closure func(*models.Post, types.Encoder)) _PostPropertyImpl {
	return _PostPropertyImpl{
		exporter.NewPropertyWithKind(name, kind),
		closure,
	}
}
