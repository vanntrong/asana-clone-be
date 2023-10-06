package project

import (
	"errors"

	"github.com/vanntrong/asana-clone-be/entities"
	"github.com/vanntrong/asana-clone-be/user"
)

type IProjectService interface {
	Create(payload CreateProjectValidation, authorId string) (project *entities.Project, err error)
	GetById(projectId string) (project *entities.Project, err error)
	GetListMember(projectId string) (members *[]entities.ProjectUsers, err error)
	AddMember(projectId string, payload AddMemberValidation, requestId string) (err error)
	RemoveMember(projectId string, payload RemoveMemberValidation, requestId string) (err error)
	GetMyProjects(userId string) (*[]entities.Project, error)
	FindMembers(projectId string, payload FindMembersValidation) (members *[]entities.User, err error)
	FindMember(projectId string, userId string) (*entities.ProjectUsers, error)
}

type ProjectService struct {
	projectRepository IProjectRepository
	userService       user.IUserService
}

func NewProjectService(projectRepository IProjectRepository, userService user.IUserService) *ProjectService {
	return &ProjectService{projectRepository, userService}
}

func (service *ProjectService) CheckListUserValid(listUser []string) bool {
	for _, userId := range listUser {
		user, err := service.userService.GetById(userId)

		if err != nil || user == nil {
			return false
		}
	}
	return true
}

func (service *ProjectService) Create(payload CreateProjectValidation, authorId string) (project *entities.Project, err error) {
	if !service.CheckListUserValid(payload.Managers) {
		return nil, errors.New("some managers are not found")
	}

	if !service.CheckListUserValid(payload.Members) {
		return nil, errors.New("some members are not found")
	}

	project, err = service.projectRepository.Create(payload.Name, authorId, payload.Managers, payload.Members)
	return project, err
}

func (service *ProjectService) GetById(projectId string) (project *entities.Project, err error) {
	project, err = service.projectRepository.GetById(projectId)
	return project, err
}

func (service *ProjectService) GetListMember(projectId string) (members *[]entities.ProjectUsers, err error) {
	members, err = service.projectRepository.GetListMember(projectId)
	return members, err
}

func (service *ProjectService) AddMember(projectId string, payload AddMemberValidation, requestId string) (err error) {
	if !service.CheckListUserValid(payload.Members) {
		return errors.New("some members are not found")
	}

	project, err := service.projectRepository.GetById(projectId)

	if err != nil || project == nil {
		return err
	}

	err = service.projectRepository.AddMember(projectId, payload.Members)
	return err
}

func (service *ProjectService) RemoveMember(projectId string, payload RemoveMemberValidation, requestId string) (err error) {
	if !service.CheckListUserValid(payload.Members) {
		return errors.New("ome members are not found")
	}

	project, err := service.projectRepository.GetById(projectId)

	if err != nil || project == nil {
		return err
	}

	err = service.projectRepository.RemoveMember(projectId, payload.Members)
	return err
}

func (service *ProjectService) GetMyProjects(userId string) (*[]entities.Project, error) {
	projects, err := service.projectRepository.GetMyProjects(userId)
	return projects, err
}

func (service *ProjectService) FindMembers(projectId string, payload FindMembersValidation) (members *[]entities.User, err error) {
	members, err = service.projectRepository.FindMembers(projectId, payload)
	return members, err
}

func (service *ProjectService) FindMember(projectId string, userId string) (*entities.ProjectUsers, error) {
	member, err := service.projectRepository.FindMember(projectId, userId)
	return member, err
}
