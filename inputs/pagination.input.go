package inputs

const DEFAULT_PAGE_SIZE = 10

type PaginationInput struct {
	PageNumber int `form:"pageNumber" binding:"omitempty,min=1"`
	PageSize   int `form:"pageSize" binding:"omitempty,min=1"`
}

func (p *PaginationInput) Limit() int {
	if p.PageSize == 0 {
		return DEFAULT_PAGE_SIZE
	}

	return p.PageSize
}

func (p *PaginationInput) Offset() int {
	if p.PageNumber == 0 {
		return 0
	}

	return (p.PageNumber - 1) * p.Limit()
}
 