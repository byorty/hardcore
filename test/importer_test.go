package test

import (
	"testing"
	"github.com/byorty/hardcore/decoder"
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/test/importers"
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
