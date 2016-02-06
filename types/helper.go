package types

type Paginator interface {
	SetPageQueryName(string)
	SetLimitQueryName(string)
	SetPage(int)
	SetLimit(int)
}

type Sorter interface {

}
