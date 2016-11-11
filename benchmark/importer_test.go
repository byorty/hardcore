package benchmark

import (
	"encoding/json"
	"github.com/byorty/hardcore/decoder"
	"github.com/byorty/hardcore/test/importers"
	"github.com/byorty/hardcore/test/models"
	"testing"
)

type ImpPost2 struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var (
	impJsonBytes = []byte(`{"id":1,"name":"name","description":"description"}`)
)

func BenchmarkJsonUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		post := new(ImpPost2)
		json.Unmarshal(impJsonBytes, post)
	}
}

func BenchmarkImporter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		post := new(models.Post)
		decoder := decoder.NewJson(impJsonBytes)
		decoder.Decode(importers.NewPost(post))
	}
}
