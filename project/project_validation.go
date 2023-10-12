package project

import "github.com/vanntrong/asana-clone-be/common"

type CreateProjectValidation struct {
	Name     string   `form:"email" json:"name" validate:"required"`
	Managers []string `form:"managers" json:"managers" validate:"required,dive,uuid"`
	Members  []string `form:"members" json:"members" validate:"required,dive,uuid"`
}

type AddMemberValidation struct {
	Members []string `form:"members" json:"members" validate:"required,dive,uuid"`
}

type RemoveMemberValidation struct {
	Members []string `form:"members" json:"members" validate:"required,dive,uuid"`
}

type FindMembersValidation struct {
	common.Pagination
}
