package odata

type Pagination struct {
	Top  *Top  `schema:"$top,omitempty"`
	Skip *Skip `schema:"$skip,omitempty"`
}

func NewPagination() Pagination {
	return Pagination{
		Top:  NewTop(),
		Skip: NewSkip(),
	}
}
