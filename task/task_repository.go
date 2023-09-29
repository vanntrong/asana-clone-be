package task

import (
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
	PatchUpdateTask(taskId string, payload *PatchUpdateTaskValidation) error
	GetListTask(query GetListTaskValidation) ([]*entities.Task, int64, error)
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (repo *TaskRepository) Create(payload *CreateTaskValidation, authorId string) (task *entities.Task, err error) {

	dueDate, err := utils.FormatTime(payload.DueDate)
	startDate, err := utils.FormatTime(payload.StartDate)

	if err != nil {
		return
	}

	// parentTaskId := uuid.Nil

	// if payload.ParentTaskId != "" {
	// 	parentTaskId = uuid.MustParse(payload.ParentTaskId)
	// }

	count, err := repo.Count(CountTaskValidation{
		ProjectId: payload.ProjectId,
		SectionId: payload.SectionId,
	})

	if err != nil {
		return
	}

	task = &entities.Task{
		Title:        payload.Title,
		Description:  payload.Description,
		StartDate:    startDate,
		DueDate:      dueDate,
		IsDone:       payload.IsDone,
		Tags:         payload.Tags,
		SectionId:    uuid.MustParse(payload.SectionId),
		AssigneeId:   uuid.MustParse(payload.AssigneeId),
		ProjectId:    uuid.MustParse(payload.ProjectId),
		CreatedById:  uuid.MustParse(authorId),
		Order:        count + 1,
		ParentTaskId: payload.ParentTaskId,
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
	dueDate, err := utils.FormatTime(payload.DueDate)
	startDate, err := utils.FormatTime(payload.StartDate)

	if err != nil {
		return
	}

	task = &entities.Task{
		Title:        payload.Title,
		Description:  payload.Description,
		StartDate:    startDate,
		DueDate:      dueDate,
		IsDone:       payload.IsDone,
		Tags:         payload.Tags,
		SectionId:    uuid.MustParse(payload.SectionId),
		AssigneeId:   uuid.MustParse(payload.AssigneeId),
		ParentTaskId: payload.ParentTaskId,
	}

	err = repo.db.Model(&task).Clauses(clause.Returning{}).Where("id = ?", taskId).Updates(task).Error

	return
}

func (repo *TaskRepository) PatchUpdateTask(taskId string, payload *PatchUpdateTaskValidation) (err error) {
	// dueDate, err := utils.FormatTime(payload.DueDate)
	// startDate, err := utils.FormatTime(payload.StartDate)

	// if err != nil {
	// 	return
	// }

	task := &entities.Task{
		Title:       payload.Title,
		Description: payload.Description,
		// StartDate:    startDate,
		// DueDate:      dueDate,
		IsDone:       payload.IsDone,
		Tags:         payload.Tags,
		SectionId:    uuid.MustParse(payload.SectionId),
		AssigneeId:   uuid.MustParse(payload.AssigneeId),
		ParentTaskId: payload.ParentTaskId,
	}

	err = repo.db.Model(&task).Clauses(clause.Returning{}).Where("id = ?", taskId).Updates(task).Error

	return
}

func (repo *TaskRepository) GetListTask(query GetListTaskValidation) (tasks []*entities.Task, total int64, err error) {
	skip := utils.GetSkipValue(query.Page, query.Limit)

	err = repo.db.Model(&tasks).Preload("Assignee").Preload("CreatedBy").Preload("ParentTask").
		Where("project_id = ?", query.ProjectId).Where("is_deleted = ?", false).
		Where("section_id = ?", query.SectionId).
		Order("tasks.order asc").
		Limit(query.Limit).
		Offset(skip).
		Count(&total).
		Find(&tasks).Error

	return
}

func (repo *TaskRepository) Count(query CountTaskValidation) (count int64, err error) {
	err = repo.db.Model(&entities.Task{}).
		Where("project_id = ?", query.ProjectId).
		Where("section_id = ?", query.SectionId).
		Count(&count).Error
	return
}
