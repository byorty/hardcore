package test

import (
	"github.com/byorty/hardcore/decoder"
	"github.com/byorty/hardcore/test/importers"
	"github.com/byorty/hardcore/test/models"
	"testing"
)

func TestImporter(t *testing.T) {
	post := new(models.Post)
	data := []byte(`{"id":1,"name":"name","description":"description"}`)
	decoder := decoder.NewJson()
	decoder.Decode(data, importers.NewPost(post))
	if post.GetId() != 1 || post.GetName() != "name" || post.GetDescription() != "description" {
		t.Fail()
	}
}
