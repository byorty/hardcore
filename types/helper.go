package types

type Paginator interface {
	GetPage() int
	SetPage(int)
	GetItemsPerPage() int
	SetItemsPerPage(int)
	GetOffset() int
	SetCount(int)
	GetPageCount() int
}

type Sorter interface {

}
