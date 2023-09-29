package task

import (
	"github.com/vanntrong/asana-clone-be/common"
)

type CreateTaskValidation struct {
	Title        string `form:"title" json:"title" validate:"required"`
	Description  string `form:"description" json:"description"`
	StartDate    string `form:"start_date" json:"start_date"`
	DueDate      string `form:"due_date" json:"due_date" validate:"required"`
	AssigneeId   string `form:"assignee_id" json:"assignee_id" validate:"required,uuid"`
	ProjectId    string `form:"project_id" json:"project_id" validate:"required,uuid"`
	SectionId    string `form:"section_id" json:"section_id" validate:"required,uuid"`
	IsDone       bool   `form:"is_done" json:"is_done" validate:"boolean"`
	Tags         string `form:"tags" json:"tags"`
	ParentTaskId string `form:"parent_task_id" json:"parent_task_id" validate:"omitempty,uuid"`
}

type UpdateTaskValidation struct {
	Title        string `form:"title" json:"title" validate:"required"`
	Description  string `form:"description" json:"description" validate:"required"`
	StartDate    string `form:"start_date" json:"start_date"`
	DueDate      string `form:"due_date" json:"due_date" validate:"required"`
	AssigneeId   string `form:"assignee_id" json:"assignee_id" validate:"required,uuid"`
	IsDone       bool   `form:"is_done" json:"is_done" validate:"required,boolean"`
	SectionId    string `form:"section_id" json:"section_id" validate:"required,uuid"`
	Tags         string `form:"tags" json:"tags"`
	ParentTaskId string `form:"parent_task_id" json:"parent_task_id" validate:"omitempty,uuid"`
}

type GetListTaskValidation struct {
	common.Pagination
	ProjectId string `form:"project_id" json:"project_id" validate:"required,uuid"`
	SectionId string `form:"section_id" json:"section_id" validate:"required,uuid"`
}

type CountTaskValidation struct {
	ProjectId string `form:"project_id" json:"project_id" validate:"required,uuid"`
	SectionId string `form:"section_id" json:"section_id" validate:"required,uuid"`
}
