package api

import (
    "github.com/byorty/hardcore/types"
    "github.com/byorty/hardcore/test/models"
)

type Post struct { 
}

func NewPost() types.ActionController {
	// implement me
    return nil
}


func (p *Post) List(page int, search string) {

}

func (p *Post) View(post *models.Post) {

}

func (p *Post) Edit(form types.Form, post *models.Post, name string, description string) {

}

