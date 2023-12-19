package appdb

import (
	"math"
)

type Paginator struct {
	Limit  int
	Offset int
}

func NewPaginator(totalAmount, limit, page int) *Paginator {
	if page < 1 {
		page = 1
	}

	totalPage := int(math.Ceil(float64(totalAmount) / float64(limit)))
	if page > totalPage {
		page = totalPage
	}

	offset := limit * (page - 1)

	return &Paginator{
		Limit:  limit,
		Offset: offset,
	}
}
