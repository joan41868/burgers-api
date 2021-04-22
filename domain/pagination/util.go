package pagination

import (
	"net/http"
	"strconv"
)

const (
	DEFAULT_PAGE_SIZE int = 100
	MAX_PAGE_SIZE     int = 1000
)

func GetPaginatedListFromRequest(r * http.Request, count int) *PaginatedList {
	page := parseInt(r.URL.Query().Get("pageNum"), 1)
	perPage := parseInt(r.URL.Query().Get("perPage"), DEFAULT_PAGE_SIZE)
	if perPage <= 0 {
		perPage = DEFAULT_PAGE_SIZE
	}
	if perPage > MAX_PAGE_SIZE {
		perPage = MAX_PAGE_SIZE
	}
	return NewPaginatedList(page, perPage, count)
}

func parseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}
