package project

import (
	"github.com/google/uuid"
	"github.com/vanntrong/asana-clone-be/entities"
	"gorm.io/gorm"
)

type IProjectRepository interface {
	Create(name string, authorId string, managers []string) (*entities.Project, error)
	GetById(projectId string) (*entities.Project, error)
	GetListMember(projectId string) (*[]entities.ProjectUsers, error)
	AddMember(projectId string, members []string) error
	RemoveMember(projectId string, members []string) error
}

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db}
}

func (repo *ProjectRepository) Create(name string, authorId string, managers []string) (*entities.Project, error) {
	project := &entities.Project{
		Name:        name,
		CreatedById: uuid.MustParse(authorId),
	}

	err := repo.db.Create(project).Error

	if err != nil {
		return nil, err
	}

	projectUsers := []entities.ProjectUsers{}

	for _, managerId := range managers {
		projectUsers = append(projectUsers, entities.ProjectUsers{
			UserId:    uuid.MustParse(managerId),
			ProjectId: project.ID,
			Role:      "manager",
		})
	}

	err = repo.db.Create(&projectUsers).Error

	return project, err
}

func (repo *ProjectRepository) GetById(projectId string) (*entities.Project, error) {
	project := &entities.Project{}
	result := repo.db.Where("id = ?", projectId).Preload("CreatedBy").Preload("Users").First(project)
	err := result.Error

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return project, err
}

func (repo *ProjectRepository) GetListMember(projectId string) (members *[]entities.ProjectUsers, err error) {
	members = &[]entities.ProjectUsers{}
	err = repo.db.Where("project_id = ?", projectId).Find(members).Error

	return
}

func (repo *ProjectRepository) AddMember(projectId string, members []string) error {
	projectUsers := []entities.ProjectUsers{}

	for _, memberId := range members {
		projectUsers = append(projectUsers, entities.ProjectUsers{
			UserId:    uuid.MustParse(memberId),
			ProjectId: uuid.MustParse(projectId),
		})
	}

	repo.db.Create(&projectUsers)

	return nil
}

func (repo *ProjectRepository) RemoveMember(projectId string, members []string) error {
	var err error
	for _, memberId := range members {
		err = repo.db.Delete(&entities.ProjectUsers{}, "user_id = ? AND project_id = ?", memberId, projectId).Error
	}

	return err
}
