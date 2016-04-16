package helper

import (
	"github.com/byorty/hardcore/types"
	"math"
)

type PaginatorImpl struct {
	page         int
	itemsPerPage int
	count        int
}

func NewPaginator() types.Paginator {
	return &PaginatorImpl{
		page:         1,
		itemsPerPage: 20,
	}
}

func (p PaginatorImpl) GetPage() int {
	return p.page
}

func (p *PaginatorImpl) SetPage(page int) {
	p.page = page
}

func (p PaginatorImpl) GetItemsPerPage() int {
	return p.itemsPerPage
}

func (p *PaginatorImpl) SetItemsPerPage(itemsPerPage int) {
	p.itemsPerPage = itemsPerPage
}

func (p PaginatorImpl) GetOffset() int {
	return (p.page - 1) * p.itemsPerPage
}

func (p *PaginatorImpl) SetCount(count int) {
	p.count = count
}

func (p PaginatorImpl) GetPageCount() int {
	return int(math.Ceil(float64(p.count) / float64(p.itemsPerPage)))
}
