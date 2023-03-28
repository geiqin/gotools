package paginate

import (
	"github.com/geiqin/gotools/helper"
)

type Paginator struct {
	Paged     int64 `json:"paged,omitempty"`
	Total     int64 `json:"total,omitempty"`
	PageCount int64 `json:"page_count,omitempty"`
	PageSize  int64 `json:"page_size,omitempty"`
	PrevPage  int64 `json:"prev_page,omitempty"`
	LastPage  int64 `json:"last_page,omitempty"`
}

func New(paged int32, pageSize ...int32) *Paginator {
	p, _ := helper.ToInt64(paged)
	if pageSize != nil {
		s, _ := helper.ToInt64(pageSize[0])
		return NewFromInt(p, s)
	}
	return NewFromInt(p)
}

func NewFromInt(paged int64, pageSize ...int64) *Paginator {
	var psize int64
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
	return helper.StringToInt(helper.ToString(offset))
}

func (a *Paginator) Limit() int {
	return helper.StringToInt(helper.ToString(a.PageSize))
}

func (a *Paginator) calculate() {
	a.PageCount = (a.Total + a.PageSize - 1) / a.PageSize
	a.LastPage = a.Paged + 1
	a.PrevPage = a.Paged - 1
	if a.LastPage > a.PageCount {
		a.LastPage = a.PageCount
	}
	if a.PrevPage < 1 {
		a.PrevPage = 1
	}
}

func (a *Paginator) ToPager(pbPager interface{}) *interface{} {
	a.calculate()
	helper.ConvertData(pbPager, a)
	/*
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
	*/
	return &pbPager
}

func Top(top int) int {
	if top > 0 {
		return top
	}
	return 20
}
