package test

import (
	"github.com/byorty/hardcore/decoder"
	"github.com/byorty/hardcore/test/importers"
	"github.com/byorty/hardcore/test/models"
	"testing"
	"fmt"
	"github.com/byorty/hardcore/slice"
)

func TestImporter(t *testing.T) {
	user := new(models.User)
	data := []byte(`{"id":1, "postIds": [1,2,3,4,5],"email":"name@mail.com", "bool": true, "posts": [{"id": 1, "name": "name#1"}, {"id": 2, "name": "name#2"}]}`)
	decoder := decoder.NewJson(data)
	decoder.Decode(importers.NewUser(user))
	fmt.Println(user)
	fmt.Println(user.GetPostIds())
	fmt.Println(user.GetPosts()[1])

	intSlice := slice.NewInt64()
	intSlice.Add(1)
	intSlice.Add(2)
	intSlice.Add(3)
	intSlice.Add(4)
	intSlice.Add(5)

	if user.GetId() != 1 ||
		user.GetEmail() != "name@mail.com" ||
		user.GetPostIds().Get(0) != intSlice.Get(0) ||
		user.GetPosts().Get(0).GetName() != "name#1" {
		t.Fail()
	}
}
