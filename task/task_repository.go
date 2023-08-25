package task

import (
	"time"

	"github.com/google/uuid"
	"github.com/vanntrong/asana-clone-be/entities"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	Create(payload *CreateTaskValidation, authorId string) (*entities.Task, error)
	GetById(taskId string) (*entities.Task, error)
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (repo *TaskRepository) Create(payload *CreateTaskValidation, authorId string) (*entities.Task, error) {
	dueDate, err := time.Parse(time.RFC3339, payload.DueDate)

	if err != nil {
		return nil, err
	}

	parentTaskId := uuid.Nil

	if payload.ParentTaskId != "" {
		parentTaskId = uuid.MustParse(payload.ParentTaskId)
	}

	task := &entities.Task{
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

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (repo *TaskRepository) GetById(taskId string) (*entities.Task, error) {
	task := &entities.Task{}
	result := repo.db.Where("id = ?", taskId).Preload("Assignee").Preload("CreatedBy").Preload("ParentTask").First(task)

	if result.Error != nil {
		return nil, result.Error
	}

	return task, nil
}
