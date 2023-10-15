package tags

import (
	"github.com/google/uuid"
	"github.com/vanntrong/asana-clone-be/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITagsRepository interface {
	CreateTag(payload CreateTagValidation) (tag *entities.Tag, err error)
	UpdateTag(tagId string, payload UpdateTagValidation) (tag *entities.Tag, err error)
	AddTagToTask(payload AddTagToTaskValidation) (taskTag *entities.TaskTags, err error)
	RemoveTagFromTask(payload RemoveTagFromTaskValidation) (err error)
	FindTag(tagId string) (tag *entities.Tag, err error)
	FindTags(projectId string) (tags []*entities.Tag, err error)
}

type TagsRepository struct {
	db *gorm.DB
}

func NewTagsRepository(db *gorm.DB) *TagsRepository {
	return &TagsRepository{
		db: db,
	}
}

func (r *TagsRepository) CreateTag(payload CreateTagValidation) (tag *entities.Tag, err error) {
	projectId := uuid.MustParse(payload.ProjectId)

	tag = &entities.Tag{
		Name:      payload.Name,
		Color:     payload.Color,
		ProjectId: projectId,
	}

	err = r.db.Create(tag).Error

	return
}

func (r *TagsRepository) UpdateTag(tagId string, payload UpdateTagValidation) (tag *entities.Tag, err error) {
	tag = &entities.Tag{
		Name:  payload.Name,
		Color: payload.Color,
	}

	err = r.db.Model(&tag).Clauses(clause.Returning{}).Where("id = ?", tagId).Updates(&payload).Error

	return
}

func (r *TagsRepository) FindTag(tagId string) (tag *entities.Tag, err error) {
	tag = &entities.Tag{}

	err = r.db.Where("id = ?", tagId).Preload("Project").First(tag).Error

	return
}

func (r *TagsRepository) FindTags(projectId string) (tags []*entities.Tag, err error) {
	err = r.db.Where("project_id = ?", projectId).Find(&tags).Error

	return
}

func (r *TagsRepository) AddTagToTask(payload AddTagToTaskValidation) (taskTag *entities.TaskTags, err error) {
	tagId := uuid.MustParse(payload.TagId)
	taskId := uuid.MustParse(payload.TaskId)

	taskTag = &entities.TaskTags{
		TaskId: taskId,
		TagId:  tagId,
	}

	err = r.db.Create(taskTag).Error

	return
}

func (r *TagsRepository) RemoveTagFromTask(payload RemoveTagFromTaskValidation) (err error) {
	tagId := uuid.MustParse(payload.TagId)
	taskId := uuid.MustParse(payload.TaskId)

	err = r.db.Where("task_id = ? AND tag_id = ?", taskId, tagId).Delete(&entities.TaskTags{}).Error

	return
}
