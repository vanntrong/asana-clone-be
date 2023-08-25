package project

import "github.com/vanntrong/asana-clone-be/entities"

type Role string

const (
	Manager Role = "manager"
	Member  Role = "member"
)

func IsUserExistInRole(project *entities.Project, userId string, role Role) bool {
	var list []entities.User

	if role == Manager {
		list = project.Managers
	} else {
		list = project.Users
	}

	for _, user := range list {
		if user.ID.String() == userId {
			return true
		}
	}

	return false
}
