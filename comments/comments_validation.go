package comments

import "github.com/vanntrong/asana-clone-be/common"

type CreateCommentValidation struct {
	Content string `json:"content" form:"content" validate:"required"`
	TaskId  string `json:"task_id" form:"task_id" validate:"required,uuid"`
}

type GetListByTaskIdValidation struct {
	common.Pagination
	TaskId string `json:"taskId" form:"task_id" validate:"required,uuid"`
}

type UpdateCommentValidation struct {
	Content string `json:"content" form:"content" validate:"required"`
}
