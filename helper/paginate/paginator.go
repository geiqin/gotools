package paginate

import (
	"github.com/geiqin/gotools/helper"
)

type Paginator struct {
	Paged     int64 `json:"paged"`
	Total     int64 `json:"total"`
	PageCount int64 `json:"page_count"`
	PageSize  int64 `json:"page_size"`
	PrevPage  int64 `json:"prev_page"`
	LastPage  int64 `json:"last_page"`
}

func New(paged int32, pageSize ...int32) *Paginator {
	entity := &Paginator{}
	entity.Paged, _ = helper.ToInt64(paged)
	if pageSize != nil {
		s, _ := helper.ToInt64(pageSize[0])
		entity.PageSize = s
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
	if a.Paged < 1 {
		a.Paged = 1
	}
	if a.PageSize <= 0 {
		a.PageSize = 20
	}
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
	return &pbPager
}

func Top(top int) int {
	if top > 0 {
		return top
	}
	return 20
}
