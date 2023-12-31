package project

import (
	"github.com/google/uuid"
	"github.com/vanntrong/asana-clone-be/entities"
	"github.com/vanntrong/asana-clone-be/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IProjectRepository interface {
	Create(name string, authorId string, managers []string, members []string) (*entities.Project, error)
	GetMyProjects(userId string) (*[]entities.Project, error)
	GetById(projectId string) (*entities.Project, error)
	GetListMember(projectId string) (*[]entities.ProjectUsers, error)
	AddMember(projectId string, members []string) error
	RemoveMember(projectId string, members []string) error
	FindMembers(projectId string, payload FindMembersValidation) (members *[]entities.User, total int64, err error)
	FindMember(projectId string, userId string) (*entities.ProjectUsers, error)
}

type ProjectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{db}
}

func (repo *ProjectRepository) Create(name string, authorId string, managers []string, members []string) (*entities.Project, error) {
	project := &entities.Project{
		Name:        name,
		CreatedById: uuid.MustParse(authorId),
	}

	err := repo.db.Clauses(clause.Returning{}).Create(project).Preload("CreatedBy").Error

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

	for _, memberId := range members {
		projectUsers = append(projectUsers, entities.ProjectUsers{
			UserId:    uuid.MustParse(memberId),
			ProjectId: project.ID,
			Role:      "member",
		})
	}

	err = repo.db.Create(&projectUsers).Error

	if err != nil {
		return nil, err
	}

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

func (repo *ProjectRepository) GetMyProjects(userId string) (*[]entities.Project, error) {
	projects := &[]entities.Project{}
	err := repo.db.Raw(`
	select * from projects p
	join project_users pu on p.id = pu.project_id
	where pu.user_id = ?
	`, userId).Scan(projects).Error

	return projects, err
}

func (repo *ProjectRepository) FindMembers(projectId string, payload FindMembersValidation) (members *[]entities.User, total int64, err error) {
	members = &[]entities.User{}
	skip := utils.GetSkipValue(payload.Page, payload.Limit)

	if err != nil {
		return nil, 0, err
	}

	stringQuery := `
	select u.id,u."name",u.email,u.avatar from users u 
	join project_users pu ON pu.user_id = u.id 
	where pu.project_id = ? and u.is_active = true and u.deleted_at is null
	`

	stringQueryCount := `
	select count(u.id) as count from users u 
	join project_users pu ON pu.user_id = u.id 
	where pu.project_id = ? and u.is_active = true and u.deleted_at is null
	`

	if payload.Keyword != "" {
		stringQuery += " and (u.name ilike '%" + payload.Keyword + "%' or u.email ilike '%" + payload.Keyword + "%')"
		stringQueryCount += " and (u.name ilike '%" + payload.Keyword + "%' or u.email ilike '%" + payload.Keyword + "%')"
	}

	stringQuery += " limit " + utils.ConvertIntToString(payload.Limit) + " offset " + utils.ConvertIntToString(skip)

	err = repo.db.Raw(stringQuery, projectId).Scan(members).Error
	err = repo.db.Raw(stringQueryCount, projectId).Scan(&total).Error

	return
}

func (repo *ProjectRepository) FindMember(projectId string, userId string) (*entities.ProjectUsers, error) {
	projectUser := &entities.ProjectUsers{}
	result := repo.db.
		Where("project_id = ?", projectId).
		Where("user_id = ?", userId).
		First(projectUser)
	err := result.Error

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return projectUser, err
}
