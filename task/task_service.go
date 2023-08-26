package task

import (
	"errors"

	"github.com/vanntrong/asana-clone-be/entities"
	"github.com/vanntrong/asana-clone-be/project"
)

type ITaskService interface {
	Create(payload *CreateTaskValidation, authorId string) (*entities.Task, error)
	GetById(taskId string, userId string) (*entities.Task, error)
}

type TaskService struct {
	taskRepository ITaskRepository
	projectService project.IProjectService
}

func NewTaskService(taskRepository ITaskRepository, projectService project.IProjectService) *TaskService {
	return &TaskService{
		taskRepository,
		projectService}
}

func (service *TaskService) Create(payload *CreateTaskValidation, authorId string) (*entities.Task, error) {
	projectFound, err := service.projectService.GetById(payload.ProjectId)

	if err != nil || projectFound == nil {
		return nil, errors.New("project not found")
	}

	task, err := service.taskRepository.Create(payload, authorId)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (service *TaskService) GetById(taskId string, userId string) (*entities.Task, error) {
	task, err := service.taskRepository.GetById(taskId)

	if err != nil {
		return nil, err
	}

	return task, nil
}
