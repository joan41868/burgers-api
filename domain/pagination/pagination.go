package pagination

type PaginatedList struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	PageCount  int         `json:"page_count"`
	TotalCount int         `json:"total_count"`
	Items      interface{} `json:"items"`
}


// NewPaginatedList creates a new PaginatedList
func NewPaginatedList(page, perPage, total int) *PaginatedList {
	if perPage < 1 {
		perPage = 15
	}
	pageCount := -1
	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}
	if page < 1 {
		page = 1
	}

	return &PaginatedList{
		Page:       page,
		PerPage:    perPage,
		TotalCount: total,
		PageCount:  pageCount,
	}
}

// Offset - offset to be used in sql statement
func (p *PaginatedList) Offset() int {
	return (p.Page - 1) * p.PerPage
}

// Limit - limit to be used in SQL statement
func (p *PaginatedList) Limit() int {
	return p.PerPage
}
