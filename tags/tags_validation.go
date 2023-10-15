package tags

type CreateTagValidation struct {
	Name      string `form:"name" json:"name" validate:"required"`
	Color     string `gorm:"not null;index" json:"color" validate:"required,hexcolor"`
	ProjectId string `form:"project_id" json:"project_id" validate:"required,uuid"`
}

type UpdateTagValidation struct {
	Name  string `form:"name" json:"name" validate:"required"`
	Color string `gorm:"not null;index" json:"color" validate:"required,hexcolor"`
}

type AddTagToTaskValidation struct {
	TagId  string `form:"tag_id" json:"tag_id" validate:"required,uuid"`
	TaskId string `form:"task_id" json:"task_id" validate:"required,uuid"`
}

type RemoveTagFromTaskValidation struct {
	TagId  string `form:"tag_id" json:"tag_id" validate:"required,uuid"`
	TaskId string `form:"task_id" json:"task_id" validate:"required,uuid"`
}

type GetListTagsValidation struct {
	ProjectId string `form:"project_id" json:"project_id" validate:"required,uuid"`
}
