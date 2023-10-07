package user

import (
	"github.com/vanntrong/asana-clone-be/common"
)

type GetListUserQuery struct {
	common.Pagination
	ExcludeInProject string `form:"exclude_in_project" validate:"omitempty,uuid"`
}

type CreateUserValidation struct {
	Email    string `form:"email" json:"email" validate:"required,email"`
	Name     string `form:"name" json:"name" validate:"required"`
	Provider string `form:"provider" json:"provider" validate:"omitempty"`
	Avatar   string `form:"avatar" json:"avatar" validate:"omitempty"`
	IsActive bool   `form:"is_active" json:"is_active" validate:"omitempty,bool"`
}
