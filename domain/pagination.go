package domain

import (
	"math"
	"net/http"
	"strconv"
)

type Pagination struct {
	TotalPages *int `json:"total_pages,omitempty"`
	TotalCount *int `json:"total_count,omitempty"`
	Page       *int `json:"page,omitempty"`
	Limit      *int `json:"limit,omitempty"`
}

type FilterPagination struct {
	Limit   int    `json:"limit,omitempty"`
	Page    int    `json:"page,omitempty"`
	OrderBy string `json:"orderBy,omitempty"`
}

const defaultSize = 10

func (data *FilterPagination) SetLimit(limitQuery string) error {
	if limitQuery == "" {
		data.Limit = defaultSize
		return nil
	}
	n, err := strconv.Atoi(limitQuery)
	if err != nil {
		return err
	}
	data.Limit = n

	return nil
}

func (data *FilterPagination) SetPage(pageQuery string) error {
	if pageQuery == "" {
		data.Page = 1
		return nil
	}

	n, err := strconv.Atoi(pageQuery)
	if err != nil {
		return err
	}

	if n == 0 {
		data.Page = 1
	}

	data.Page = n

	return nil
}

func (data *FilterPagination) SetOrderBy(orderByQuery string) {
	data.OrderBy = orderByQuery
}

func (data *FilterPagination) GetOffset() int {
	if data.Page == 0 {
		return 0
	}
	return (data.Page - 1) * data.Limit
}

func (data *FilterPagination) GetLimit() int {
	return data.Limit
}

func (data *FilterPagination) GetOrderBy() string {
	return data.OrderBy
}

func (data *FilterPagination) GetPage() int {
	return data.Page
}

func (data *FilterPagination) GetTotalPages(totalCount int) int {
	return int(math.Ceil(float64(totalCount) / float64(data.GetLimit())))
}

// pagination := &httpresp.Pagination{
// 	TotalCount: int(count),
// 	TotalPages: utils.GetTotalPages(int(count), pq.GetLimit()),
// 	Page:       pq.GetPage(),
// 	Limit:      pq.GetLimit(),
// }

func NewPagination(totalData int, pagination FilterPagination) Pagination {
	page := pagination.GetPage()
	limit := pagination.GetLimit()
	totalPages := pagination.GetTotalPages(totalData)
	return Pagination{
		TotalPages: &totalPages,
		TotalCount: &totalData,
		Page:       &page,
		Limit:      &limit,
	}
}

func GetPaginationFromCtx(r *http.Request) (FilterPagination, error) {
	var q FilterPagination

	// Extract query parameters from the request
	if err := q.SetPage(r.URL.Query().Get("page")); err != nil {
		return FilterPagination{}, err
	}
	if err := q.SetLimit(r.URL.Query().Get("limit")); err != nil {
		return FilterPagination{}, err
	}
	q.SetOrderBy(r.URL.Query().Get("orderBy"))

	return q, nil
}
