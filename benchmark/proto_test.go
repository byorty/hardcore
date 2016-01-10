package benchmark

import (
	"testing"
	"reflect"
	"github.com/byorty/hardcore/test/models"
)

var (
	post = new(models.Post).
		SetId(1).
		SetUserId(1).
		SetName("name").
		SetDescription("description")
	idInt64 int64 = 2
	idInt64Value = reflect.ValueOf(idInt64)
	getIdMethod = (*models.Post).GetId
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
		getter := post.Proto().GetByName("id").GetGetter()
		getter(post)
	}
}

func BenchmarkProtoCallGetMethod2(b *testing.B) {
	method := post.Proto().GetByName("id").GetGetter()
	for i := 0;i < b.N;i++ {
		method(post)
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
		setter := post.Proto().GetByName("id").GetSetter()
		setter(post, idInt64)
	}
}

func BenchmarkNativeCallSetMethod(b *testing.B) {
	for i := 0;i < b.N;i++ {
		post.SetId(idInt64)
	}
}
