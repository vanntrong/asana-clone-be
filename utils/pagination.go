package utils

import "github.com/vanntrong/asana-clone-be/common"

func GetSkipValue(page int, limit int) int {
	return (page - 1) * limit
}

func GetPaginationResponse(total int64, query *common.Pagination) *common.PaginationResponse {
	return &common.PaginationResponse{
		Total:   total,
		HasNext: (query.Page)*query.Limit < int(total),
		Pagination: common.Pagination{
			Page:      query.Page,
			Limit:     query.Limit,
			SortBy:    query.SortBy,
			SortOrder: query.SortOrder,
		},
	}
}
