package benchmark

import (
	"testing"
	"github.com/byorty/hardcore/proto"
	"github.com/byorty/hardcore/types"
	"reflect"
)

type Post struct {
	AutoPost
}

type AutoPost struct {
	id int64
	userId int64
	name string
	description string
}

func (p Post) GetId() int64 {
	return p.id
}

func (p *Post) SetId(id int64) *Post {
	p.id = id
	return p
}

func (p Post) GetUserId() int64 {
	return p.userId
}

func (p *Post) SetUserId(userId int64) *Post {
	p.userId = userId
	return p
}

func (p Post) GetName() string {
	return p.name
}

func (p *Post) SetName(name string) *Post {
	p.name = name
	return p
}

func (p Post) GetDescription() string {
	return p.description
}

func (p *Post) SetDescription(description string) *Post {
	p.description = description
	return p
}

func (p *Post) Proto() types.Proto {
	return postProto
}

type PostIdSetter func(*Post, int64) *Post

func (p PostIdSetter) Call(model interface{}, id interface{}) {
	p(model.(*Post), id.(int64))
}

type PostIdGetter func(*Post) int64

func (p PostIdGetter) Call(model interface{}) interface{} {
	return p(model.(*Post))
}

type PostUserIdSetter func(*Post, int64) *Post

func (p PostUserIdSetter) Call(model interface{}, userId interface{}) {
	p(model.(*Post), userId.(int64))
}

type PostUserIdGetter func(*Post) int64

func (p PostUserIdGetter) Call(model interface{}) interface{} {
	return p(model.(*Post))
}

type PostNameSetter func(*Post, string) *Post

func (p PostNameSetter) Call(model interface{}, name interface{}) {
	p(model.(*Post), name.(string))
}

type PostNameGetter func(*Post) string

func (p PostNameGetter) Call(model interface{}) interface{} {
	return p(model.(*Post))
}

type PostDescriptionSetter func(*Post, string) *Post

func (p PostDescriptionSetter) Call(model interface{}, description interface{}) {
	p(model.(*Post), description.(string))
}

type PostDescriptionGetter func(*Post) string

func (p PostDescriptionGetter) Call(model interface{}) interface{} {
	return p(model.(*Post))
}

type PostIdGetter2 func(*Post) int64

func (p PostIdGetter2) Call(model *Post) interface{} {
	return p(model)
}

var (
	postIdSetter PostIdSetter = (*Post).SetId
	postIdGetter PostIdGetter = (*Post).GetId
	postUserIdSetter PostUserIdSetter = (*Post).SetUserId
	postUserIdGetter PostUserIdGetter = (*Post).GetUserId
	postNameSetter PostNameSetter = (*Post).SetName
	postNameGetter PostNameGetter = (*Post).GetName
	postDescriptionSetter PostDescriptionSetter = (*Post).SetDescription
	postDescriptionGetter PostDescriptionGetter = (*Post).GetDescription
	postProto = proto.New().
		Set("id", proto.NewProperty("id", types.ProtoBasicKind, types.ProtoNoneRelation, true, postIdSetter, postIdGetter)).
		Set("userId", proto.NewProperty("user_id", types.ProtoBasicKind, types.ProtoNoneRelation, true, postUserIdSetter, postUserIdGetter)).
		Set("name", proto.NewProperty("name", types.ProtoBasicKind, types.ProtoNoneRelation, true, postNameSetter, postNameGetter)).
		Set("description", proto.NewProperty("description", types.ProtoBasicKind, types.ProtoNoneRelation, true, postDescriptionSetter, postDescriptionGetter))

	post = &Post{AutoPost{1, 1, "name", "description"}}
	idInt64 int64 = 2
	idInt64Value = reflect.ValueOf(idInt64)
	getIdMethod = (*Post).GetId
	postIdGetter2 PostIdGetter2 = (*Post).GetId
)

func BenchmarkReflectCallGetMethod(b *testing.B) {
	for i := 0;i < b.N;i++ {
		reflect.ValueOf(post).MethodByName("GetId").Call([]reflect.Value{})
	}
}

func BenchmarkReflectCallGetMethod2(b *testing.B) {
	method := reflect.ValueOf(post).MethodByName("GetId")
	value := []reflect.Value{}
	for i := 0;i < b.N;i++ {
		method.Call(value)
	}
}

func BenchmarkProtoCallGetMethod(b *testing.B) {
	for i := 0;i < b.N;i++ {
		post.Proto().GetByName("id").GetGetter().Call(post)
	}
}

func BenchmarkProtoCallGetMethod2(b *testing.B) {
	method := post.Proto().GetByName("id").GetGetter()
	for i := 0;i < b.N;i++ {
		method.Call(post)
	}
}

func BenchmarkProtoCallGetMethod3(b *testing.B) {
	for i := 0;i < b.N;i++ {
		postIdGetter2.Call(post)
	}
}

func BenchmarkNativeCallGetMethod(b *testing.B) {
	for i := 0;i < b.N;i++ {
		post.GetId()
	}
}

func BenchmarkNativeCallGetMethod2(b *testing.B) {
	for i := 0;i < b.N;i++ {
		getIdMethod(post)
	}
}

func BenchmarkReflectCallSetMethod(b *testing.B) {
	for i := 0;i < b.N;i++ {
		reflect.ValueOf(post).MethodByName("SetId").Call([]reflect.Value{idInt64Value})
	}
}

func BenchmarkProtoCallSetMethod(b *testing.B) {
	for i := 0;i < b.N;i++ {
		post.Proto().GetByName("id").GetSetter().Call(post, idInt64)
	}
}

func BenchmarkNativeCallSetMethod(b *testing.B) {
	for i := 0;i < b.N;i++ {
		post.SetId(idInt64)
	}
}
