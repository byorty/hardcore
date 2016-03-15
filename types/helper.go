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

type SessionManager interface {
	Create(RequestScope) SessionScope
	Get(RequestScope) SessionScope
	Remove(RequestScope)
	SetProvider(Cache) SessionManager
}
