package view

import "github.com/byorty/hardcore/types"

type HtmlImpl struct {
	BaseImpl
	template string
}

func (h *HtmlImpl) SetTemplate(template string) types.HtmlView {
	h.template = template
}

func (h *HtmlImpl) Render() {

}
