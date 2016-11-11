package importers

import (
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

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
