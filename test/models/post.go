package models

type Post struct {
	AutoPost
}

type Posts []*Post

type PostDao struct {
	AutoPostDao
}
