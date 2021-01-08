package paginate

import (
	"github.com/geiqin/gotools/helper"
)

type Paginator struct {
	Paged     int
	Total     int64
	PageCount int
	PageSize  int
	PrevPage  int
	LastPage  int
}

type paginatorWrap struct {
	Paged     int32
	Total     int32
	PageCount int32
	PageSize  int32
	PrevPage  int32
	LastPage  int32
}



func New(paged int32, pageSize ...int32) *Paginator {
	 p :=helper.Int32ToInt(paged)
	 if pageSize !=nil {
	 	s :=helper.Int32ToInt(pageSize[0])
	 	return NewFromInt(p,s)
	 }
	 return NewFromInt(p)
}

func NewFromInt(paged int, pageSize ...int) *Paginator {
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
	w:=&paginatorWrap{Paged: helper.IntToInt32(a.Paged)}
	v :=helper.Int64ToString(a.Total)
	t := helper.StringToInt32(v)
	w.Total =t
	w.PageSize = helper.IntToInt32(a.PageSize)
	w.PageCount = (t +  w.PageSize - 1) / w.PageSize
	a.LastPage = a.Paged + 1
	a.PrevPage = a.Paged - 1
	if a.LastPage > a.PageCount {
		a.LastPage = a.PageCount
	}
	if a.PrevPage < 1 {
		a.PrevPage = 1
	}

	helper.StructCopy(pbPager, w)
	return &pbPager
}

func Top(top int) int {
	if top > 0 {
		return top
	}
	return 20
}
