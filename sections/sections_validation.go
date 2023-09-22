package sections

type GetListSectionValidation struct {
	ProjectId string `form:"project_id" json:"project_id" validate:"required,uuid"`
}

type CreateSectionValidation struct {
	ProjectId   string `form:"project_id" json:"project_id" validate:"required,uuid"`
	Name        string `form:"name" json:"name" validate:"required"`
	Description string `form:"description" json:"description" validate:"required"`
}

type UpdateSectionValidation struct {
	Name        string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`
}
