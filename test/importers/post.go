package importers

import (
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

type _PostImpl struct {
	model *models.Post
	props map[string]_PostPropertyImpl
}

func (p _PostImpl) Get(key string) (types.ImportableProperty, bool) {
	prop, ok := p.props[key]
	return prop, ok
}

func (p _PostImpl) Decode(key string, decoder types.Decoder, value []byte) {
	p.props[key].closure(p.model, decoder, value)
}

type _PostPropertyImpl struct {
	kind types.ProtoKind
	closure func(*models.Post, types.Decoder, []byte)
}

func (p _PostPropertyImpl) GetProtoKind() types.ProtoKind {
	return p.kind
}

func newPostProperty(kind types.ProtoKind, closure func(*models.Post, types.Decoder, []byte)) _PostPropertyImpl {
	return _PostPropertyImpl{
		kind,
		closure,
	}
}

func NewPost(post *models.Post) types.Importer {
	imp := new(_PostImpl)
	imp.model = post
	imp.props = _postProperties
	return imp
}

//func NewPosts(posts models.Posts) types.Importer {
//	imp := new(importer.BaseImpl)
//	imp.SetProperties(postProperties)
//	imp.SetImportable(posts)
//	return imp
//}

var (
	_postProperties = map[string]_PostPropertyImpl{
		"id": newPostProperty(types.ProtoInt64Kind, func(post *models.Post, decoder types.Decoder, value []byte) {
			post.SetId(decoder.DecodeInt64(value))
		}),
		"name": newPostProperty(types.ProtoStringKind, func(post *models.Post, decoder types.Decoder, value []byte) {
			post.SetName(decoder.DecodeString(value))
		}),
		"description": newPostProperty(types.ProtoStringKind, func(post *models.Post, decoder types.Decoder, value []byte) {
			post.SetDescription(decoder.DecodeString(value))
		}),
	}
)
