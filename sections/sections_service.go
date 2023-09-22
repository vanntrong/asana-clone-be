package sections

import (
	"fmt"

	"github.com/vanntrong/asana-clone-be/entities"
	"github.com/vanntrong/asana-clone-be/project"
)

type ISectionsService interface {
	GetList(userId string, projectId string) ([]*entities.Section, error)
	GetById(userId string, sectionId string) (*entities.Section, error)
	CreateSection(userId string, body CreateSectionValidation) (*entities.Section, error)
	UpdateSection(userId string, sectionId string, body UpdateSectionValidation) (*entities.Section, error)
}

type SectionsService struct {
	sectionRepository ISectionsRepository
	projectService    project.IProjectService
}

func NewSectionsService(sectionRepository ISectionsRepository, projectService project.IProjectService) *SectionsService {
	return &SectionsService{sectionRepository: sectionRepository, projectService: projectService}
}

func (service *SectionsService) GetList(userId string, projectId string) (sections []*entities.Section, err error) {
	projectMembers, err := service.projectService.GetListMember(projectId)

	isMember := project.IsMember(projectMembers, userId)

	if err != nil || !isMember {
		return nil, fmt.Errorf("you are not member of this project")
	}

	sections, err = service.sectionRepository.GetList(projectId)

	return
}

func (service *SectionsService) CreateSection(userId string, body CreateSectionValidation) (section *entities.Section, err error) {
	projectMembers, err := service.projectService.GetListMember(body.ProjectId)

	isMember := project.IsMember(projectMembers, userId)

	if err != nil || !isMember {
		return nil, fmt.Errorf("you are not member of this project")
	}

	section, err = service.sectionRepository.CreateSection(body)

	return
}

func (service *SectionsService) UpdateSection(userId string, sectionId string, body UpdateSectionValidation) (section *entities.Section, err error) {

	section, err = service.sectionRepository.GetById(sectionId)

	if err != nil {
		return nil, err
	}

	projectMembers, err := service.projectService.GetListMember(section.ProjectId.String())

	isMember := project.IsMember(projectMembers, userId)

	if err != nil || !isMember {
		return nil, fmt.Errorf("you are not member of this project")
	}

	section, err = service.sectionRepository.UpdateSection(sectionId, body)

	return
}

func (service *SectionsService) GetById(userId string, sectionId string) (section *entities.Section, err error) {
	section, err = service.sectionRepository.GetById(sectionId)

	if err != nil {
		return nil, fmt.Errorf("section not found")
	}

	projectMembers, err := service.projectService.GetListMember(section.ProjectId.String())

	isMember := project.IsMember(projectMembers, userId)

	if err != nil || !isMember {
		return nil, fmt.Errorf("you are not member of this project")
	}

	return
}
