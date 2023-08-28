package task

import (
	"time"

	"github.com/google/uuid"
	"github.com/vanntrong/asana-clone-be/entities"
	"github.com/vanntrong/asana-clone-be/utils"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	Create(payload *CreateTaskValidation, authorId string) (*entities.Task, error)
	GetById(taskId string) (*entities.Task, error)
	UpdateTask(taskId string, payload *UpdateTaskValidation) (*entities.Task, error)
	GetListTask(query GetListTaskValidation) ([]*entities.Task, int64, error)
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (repo *TaskRepository) Create(payload *CreateTaskValidation, authorId string) (task *entities.Task, err error) {
	dueDate, err := time.Parse(time.RFC3339, payload.DueDate)

	if err != nil {
		return
	}

	parentTaskId := uuid.Nil

	if payload.ParentTaskId != "" {
		parentTaskId = uuid.MustParse(payload.ParentTaskId)
	}

	task = &entities.Task{
		Title:        payload.Title,
		Description:  payload.Description,
		DueDate:      dueDate,
		Status:       payload.Status,
		Tags:         payload.Tags,
		AssigneeId:   uuid.MustParse(payload.AssigneeId),
		ProjectId:    uuid.MustParse(payload.ProjectId),
		CreatedById:  uuid.MustParse(authorId),
		ParentTaskId: parentTaskId,
	}

	err = repo.db.Create(task).Error
	return
}

func (repo *TaskRepository) GetById(taskId string) (task *entities.Task, err error) {
	task = &entities.Task{}
	err = repo.db.Where("id = ?", taskId).Preload("Assignee").Preload("Project").Preload("CreatedBy").Preload("ParentTask").First(task).Error
	return
}

func (repo *TaskRepository) UpdateTask(taskId string, payload *UpdateTaskValidation) (task *entities.Task, err error) {
	dueDate, err := time.Parse(time.RFC3339, payload.DueDate)

	if err != nil {
		return
	}

	parentTaskId := uuid.Nil

	if payload.ParentTaskId != "" {
		parentTaskId = uuid.MustParse(payload.ParentTaskId)
	}

	task = &entities.Task{
		Title:        payload.Title,
		Description:  payload.Description,
		DueDate:      dueDate,
		Status:       payload.Status,
		Tags:         payload.Tags,
		AssigneeId:   uuid.MustParse(payload.AssigneeId),
		ParentTaskId: parentTaskId,
	}

	err = repo.db.Model(&task).Clauses(clause.Returning{}).Where("id = ?", taskId).Updates(task).Error

	return
}

func (repo *TaskRepository) GetListTask(query GetListTaskValidation) (tasks []*entities.Task, total int64, err error) {
	skip := utils.GetSkipValue(query.Page, query.Limit)

	err = repo.db.Model(&tasks).Preload("Assignee").Preload("Project").Preload("CreatedBy").Preload("ParentTask").
		Where("project_id = ?", query.ProjectId).Where("is_deleted = ?", false).
		Limit(query.Limit).
		Offset(skip).
		Count(&total).
		Find(&tasks).Error

	return
}
