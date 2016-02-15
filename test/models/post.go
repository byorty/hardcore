package models

import (
	"github.com/byorty/hardcore/types"
)

type Post struct {
	AutoPost
}

type Posts []*Post

type PostDao struct {
	AutoPostDao
}

func (p *PostDao) Init(db types.DB) {}
