package project

import (
	"github.com/vanntrong/asana-clone-be/entities"
)

type Role string

const (
	Manager Role = "manager"
	Member  Role = "member"
)

func IsMember(members *[]entities.ProjectUsers, userId string) bool {
	for _, user := range *members {
		if user.UserId.String() == userId {
			return true
		}
	}

	return false
}

func IsManager(members *[]entities.ProjectUsers, userId string) bool {

	for _, user := range *members {
		if user.UserId.String() == userId && Role(user.Role) == Manager {
			return true
		}
	}

	return false
}
