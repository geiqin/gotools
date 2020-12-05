package paginate

import "github.com/geiqin/gotools/helper"

type Paginator struct {
	Paged     int
	Total     int
	PageCount int
	PageSize  int
	PrevPage  int
	LastPage  int
}

func New(paged int, pageSize ...int) *Paginator {
	var psize int
	if pageSize != nil {
		if pageSize[0] > 0 {
			psize = pageSize[0]
		}
	}
	if paged < 1 {
		paged = 1
	}

	if psize < 1 {
		psize = 20
	}

	entity := &Paginator{
		PageSize: psize,
		Paged:    paged,
	}
	return entity
}

func (a *Paginator) Offset() int {
	offset := (a.Paged - 1) * a.PageSize
	return offset
}

func (a *Paginator) Limit() int {
	return a.PageSize
}

func (a *Paginator) ToPager(pbPager interface{}) *interface{} {
	a.PageCount = (a.Total + a.PageSize - 1) / a.PageSize
	a.LastPage = a.Paged + 1
	a.PrevPage = a.Paged - 1
	if a.LastPage > a.PageCount {
		a.LastPage = a.PageCount
	}
	if a.PrevPage < 1 {
		a.PrevPage = 1
	}
	helper.StructCopy(pbPager, a)
	return &pbPager
}

func Top(top int) int {
	if top > 0 {
		return top
	}
	return 20
}
