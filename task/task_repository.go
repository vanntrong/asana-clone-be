package task

import (
	"fmt"

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
	GetListTask(query GetListTaskValidation, userId string) ([]*entities.Task, int64, error)
	UpdateOrderTasks(projectId string, sectionId string, tasks []string) error
	CheckLikeExist(taskId string, userId string) (isExist bool, err error)
	LikeTask(taskId string, userId string) (err error)
	UnLikeTask(taskId string, userId string) (err error)
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
	err = repo.db.Where("id = ?", taskId).Preload("Assignee").Preload("Project").Preload("CreatedBy").Preload("ParentTask").Preload("TagsList").First(task).Error
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
		SectionId:    uuid.MustParse(payload.SectionId),
		AssigneeId:   uuid.MustParse(payload.AssigneeId),
		ParentTaskId: payload.ParentTaskId,
	}

	err = repo.db.Model(&task).Clauses(clause.Returning{}).Where("id = ?", taskId).Updates(task).Error

	return
}

func (repo *TaskRepository) GetListTask(query GetListTaskValidation, userId string) (tasks []*entities.Task, total int64, err error) {
	skip := utils.GetSkipValue(query.Page, query.Limit)

	queryBuilder := repo.db.Model(&tasks).Preload("Assignee").Preload("CreatedBy").Preload("TagsList").Preload("ParentTask").
		Select("tasks.*, case when task_likes.task_id is not null then true else false end as is_liked, COALESCE(like_counts.like_count, 0) AS like_count").
		Joins("left join task_likes on tasks.id = task_likes.task_id and task_likes.user_id = ?", userId).
		Joins("left join (select task_id, count(*) AS like_count from task_likes group by task_id) as like_counts on tasks.id = like_counts.task_id").
		Joins("left join task_tags on tasks.id = task_tags.task_id").
		Where("project_id = ?", query.ProjectId).Where("deleted_at is null").
		Where("section_id = ?", query.SectionId).
		Group("tasks.id").
		Group("task_likes.task_id").Group("like_counts.like_count")
	repo.addQueryAssignee(queryBuilder, query.AssigneeIds)
	if query.IsDone {
		queryBuilder.Where("is_done = ?", query.IsDone)
	}

	if query.StartDate != "" {
		startDate, err := utils.FormatTime(query.StartDate)
		if err == nil {
			queryBuilder.Where("start_date = ?", startDate)
		}

	}

	if query.DueDate != "" {
		dueDate, err := utils.FormatTime(query.DueDate)

		if err == nil {
			queryBuilder.Where("due_date <= ?", dueDate)
		}
	}

	err = queryBuilder.Order("tasks.order asc").Limit(query.Limit).Offset(skip).Count(&total).Find(&tasks).Error

	return
}

func (repo *TaskRepository) Count(query CountTaskValidation) (count int64, err error) {
	err = repo.db.Model(&entities.Task{}).
		Where("project_id = ?", query.ProjectId).
		Where("section_id = ?", query.SectionId).
		Count(&count).Error
	return
}

func (repo *TaskRepository) UpdateOrderTasks(projectId string, sectionId string, tasks []string) error {
	for index, taskId := range tasks {
		err := repo.db.Model(&entities.Task{}).Where("id = ?", taskId).
			Where("project_id = ?", projectId).
			Where("deleted_at is null").
			Update("order", index+1).
			Update("section_id", sectionId).
			Error

		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *TaskRepository) CheckLikeExist(taskId string, userId string) (isExist bool, err error) {
	like := &entities.TaskLikes{}

	err = repo.db.Where("task_id = ?", taskId).Where("user_id = ?", userId).First(like).Error

	if err != nil || like == nil {
		return false, err
	}

	return true, nil
}

func (repo *TaskRepository) LikeTask(taskId string, userId string) (err error) {
	like := &entities.TaskLikes{
		TaskId: uuid.MustParse(taskId),
		UserId: uuid.MustParse(userId),
	}

	fmt.Print(like)

	err = repo.db.Create(like).Error

	return
}

func (repo *TaskRepository) UnLikeTask(taskId string, userId string) (err error) {
	like := &entities.TaskLikes{}

	err = repo.db.Where("task_id = ?", taskId).Where("user_id = ?", userId).Delete(like).Error

	return
}

func (repo *TaskRepository) addQueryAssignee(query *gorm.DB, assigneeIds []string) {
	if len(assigneeIds) > 0 {
		query.Where("assignee_id in (?)", assigneeIds)
	}
}
