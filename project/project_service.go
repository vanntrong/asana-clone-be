package project

import (
	"errors"

	"github.com/vanntrong/asana-clone-be/entities"
	"github.com/vanntrong/asana-clone-be/user"
)

type IProjectService interface {
	Create(payload CreateProjectValidation, authorId string) (project *entities.Project, err error)
	GetById(projectId string) (project *entities.Project, err error)
	AddMember(projectId string, payload AddMemberValidation, requestId string) (err error)
	RemoveMember(projectId string, payload RemoveMemberValidation, requestId string) (err error)
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
		return nil, errors.New("Some managers are not found")
	}

	project, err = service.projectRepository.Create(payload.Name, authorId, payload.Managers)
	return project, err
}

func (service *ProjectService) GetById(projectId string) (project *entities.Project, err error) {
	project, err = service.projectRepository.GetById(projectId)
	return project, err
}

func (service *ProjectService) AddMember(projectId string, payload AddMemberValidation, requestId string) (err error) {
	if !service.CheckListUserValid(payload.Members) {
		return errors.New("Some members are not found")
	}

	project, err := service.projectRepository.GetById(projectId)

	if err != nil || project == nil {
		return err
	}

	if !IsUserExistInRole(project, requestId, Manager) {
		return errors.New("You are not manager of this project")
	}

	err = service.projectRepository.AddMember(projectId, payload.Members)
	return err
}

func (service *ProjectService) RemoveMember(projectId string, payload RemoveMemberValidation, requestId string) (err error) {
	if !service.CheckListUserValid(payload.Members) {
		return errors.New("Some members are not found")
	}

	project, err := service.projectRepository.GetById(projectId)

	if err != nil || project == nil {
		return err
	}

	if !IsUserExistInRole(project, requestId, Manager) {
		return errors.New("You are not manager of this project")
	}

	err = service.projectRepository.RemoveMember(projectId, payload.Members)
	return err
}
