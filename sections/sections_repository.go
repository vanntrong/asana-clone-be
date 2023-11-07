package sections

import (
	"github.com/google/uuid"
	"github.com/vanntrong/asana-clone-be/entities"
	"github.com/vanntrong/asana-clone-be/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ISectionsRepository interface {
	GetList(projectId string) ([]*entities.Section, error)
	GetById(sectionId string) (*entities.Section, error)
	CreateSection(body CreateSectionValidation) (*entities.Section, error)
	UpdateSection(sectionId string, body UpdateSectionValidation) (*entities.Section, error)
	Delete(sectionId string) error
}

type SectionsRepository struct {
	db *gorm.DB
}

func NewSectionsRepository(db *gorm.DB) *SectionsRepository {
	return &SectionsRepository{db: db}
}

func (repo *SectionsRepository) GetList(projectId string) (sections []*entities.Section, err error) {
	err = repo.db.Where("project_id = ?", projectId).Where("deleted_at is null").Find(&sections).Error

	return
}

func (repo *SectionsRepository) CreateSection(body CreateSectionValidation) (section *entities.Section, err error) {
	section = &entities.Section{
		Name: body.Name,
		// Description: body.Description,
		ProjectId: uuid.MustParse(body.ProjectId),
	}

	err = repo.db.Clauses(clause.Returning{}).Create(section).Error

	return
}

func (repo *SectionsRepository) UpdateSection(sectionId string, body UpdateSectionValidation) (section *entities.Section, err error) {
	section = &entities.Section{
		Name:        body.Name,
		Description: body.Description,
	}

	err = repo.db.Clauses(clause.Returning{}).Where("id = ?", sectionId).Updates(section).Error

	return
}

func (repo *SectionsRepository) GetById(sectionId string) (section *entities.Section, err error) {
	err = repo.db.Where("id = ?", sectionId).Where("deleted_at is null").First(&section).Preload("Project").Error

	return
}

func (repo *SectionsRepository) Delete(sectionId string) error {
	err := repo.db.Model(&entities.Section{}).Where("id = ?", sectionId).Update("deleted_at", utils.GetTimeNow()).Error

	return err
}
