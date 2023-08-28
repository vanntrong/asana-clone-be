package common

type Pagination struct {
	Page      int    `form:"page,default=1" json:"page" binding:"required,min=1"`
	Limit     int    `form:"limit,default=10" json:"limit" binding:"required,min=1"`
	SortBy    string `form:"sort_by" json:"sort_by,omitempty"`
	SortOrder string `form:"sort_order" json:"sort_order,omitempty"`
	Keyword   string `form:"keyword" json:"keyword,omitempty"`
}

type PaginationResponse struct {
	Total   int64 `json:"total"`
	HasNext bool  `json:"has_next"`
	Pagination
}
