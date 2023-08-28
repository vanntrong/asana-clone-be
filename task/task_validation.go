package task

type CreateTaskValidation struct {
	Title        string `form:"title" json:"title" validate:"required"`
	Description  string `form:"description" json:"description" validate:"required"`
	DueDate      string `form:"due_date" json:"due_date" validate:"required"`
	AssigneeId   string `form:"assignee_id" json:"assignee_id" validate:"required,uuid"`
	ProjectId    string `form:"project_id" json:"project_id" validate:"required,uuid"`
	Status       string `form:"status" json:"status" validate:"required,oneof=todo in-progress in-review done"`
	Tags         string `form:"tags" json:"tags"`
	ParentTaskId string `form:"parent_task_id" json:"parent_task_id" validate:"omitempty,uuid"`
}

type UpdateTaskValidation struct {
	Title        string `form:"title" json:"title" validate:"required"`
	Description  string `form:"description" json:"description" validate:"required"`
	DueDate      string `form:"due_date" json:"due_date" validate:"required"`
	AssigneeId   string `form:"assignee_id" json:"assignee_id" validate:"required,uuid"`
	Status       string `form:"status" json:"status" validate:"omitempty,oneof=todo in-progress in-review done"`
	Tags         string `form:"tags" json:"tags"`
	ParentTaskId string `form:"parent_task_id" json:"parent_task_id" validate:"omitempty,uuid"`
}
