package api

import (
	"github.com/byorty/hardcore/test/models"
	"github.com/byorty/hardcore/types"
)

type Post struct {
}

func NewPost() types.ActionController {
	// implement me
	return nil
}

func (p *Post) List(page int, search string) types.EncodeView {

	return nil
}

func (p *Post) View(post *models.Post) types.EncodeView {

	return nil
}

func (p *Post) Edit(form types.Form, post *models.Post, name string, description string) types.EncodeView {

	return nil
}
