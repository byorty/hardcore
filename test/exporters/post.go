package exporters

import (
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

var (
	postProperties = []_PostPropertyImpl{
		newPostProperty("id", types.ProtoInt64Kind, func(post *models.Post, encoder types.Encoder) {
			encoder.EncodeInt64(post.GetId())
		}),
		newPostProperty("name", types.ProtoStringKind, func(post *models.Post, encoder types.Encoder) {
			encoder.EncodeString(post.GetName())
		}),
		newPostProperty("description", types.ProtoStringKind, func(post *models.Post, encoder types.Encoder) {
			encoder.EncodeString(post.GetDescription())
		}),
	}
)
