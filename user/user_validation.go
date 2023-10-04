package user

import (
	"github.com/vanntrong/asana-clone-be/common"
)

type GetListUserQuery struct {
	common.Pagination
	ExcludeInProject string `form:"exclude_in_project" validate:"omitempty,uuid"`
}
