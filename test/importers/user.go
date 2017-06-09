package importers

import (
	"github.com/byorty/hardcore/slice"
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

var (
	_userProperties = map[string]_UserPropertyImpl{
		"id": newUserProperty(types.ProtoInt64Kind, func(user *models.User, decoder types.Decoder, value []byte) {
			user.SetId(decoder.DecodeInt64(value))
		}),
		"email": newUserProperty(types.ProtoStringKind, func(user *models.User, decoder types.Decoder, value []byte) {
			user.SetEmail(decoder.DecodeString(value))
		}),
		"postIds": newUserProperty(types.ProtoInt64SliceKind, func(user *models.User, decoder types.Decoder, value []byte) {
			if user.GetPostIds() == nil {
				user.SetPostIds(slice.NewInt64())
			}
			user.GetPostIds().Add(decoder.DecodeInt64(value))
		}),
		"posts": newUserProperty(types.ProtoModelSliceKind, func(user *models.User, decoder types.Decoder, value []byte) {
			if !user.HasPosts() {
				user.SetPosts(make(models.Posts, 0))
			}
			post := new(models.Post)
			decoder.Decode(NewPost(post))
			user.SetPosts(append(user.GetPosts(), post))
		}),
	}
)
