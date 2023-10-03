package task

import (
	"errors"
	"reflect"

	"github.com/google/uuid"
	"github.com/vanntrong/asana-clone-be/common"
	"github.com/vanntrong/asana-clone-be/entities"
	"github.com/vanntrong/asana-clone-be/project"
	"github.com/vanntrong/asana-clone-be/utils"
)

type ITaskService interface {
	Create(payload *CreateTaskValidation, authorId string) (*entities.Task, error)
	GetById(taskId string, userId string) (*entities.Task, error)
	UpdateTask(taskId string, payload *UpdateTaskValidation, userId string) (*entities.Task, error)
	PatchUpdateTask(taskId string, payload *PatchUpdateTaskValidation, userId string) error
	DeleteTask(taskId string, userId string) error
	GetListTask(userId string, query GetListTaskValidation) ([]*entities.Task, *common.PaginationResponse, error)
	UpdateOrderTasks(projectId string, sectionId string, userId string, tasks []string) error
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

	err = service.isParentTaskValid(payload.ParentTaskId, payload.ProjectId)

	if err != nil {
		return nil, err
	}

	err = service.isAssigneeValid(payload.AssigneeId, payload.ProjectId)

	if err != nil {
		return nil, err
	}

	task, err := service.taskRepository.Create(payload, authorId)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (service *TaskService) GetById(taskId string, userId string) (*entities.Task, error) {
	return service.taskRepository.GetById(taskId)
}

func (service *TaskService) UpdateTask(taskId string, payload *UpdateTaskValidation, userId string) (*entities.Task, error) {
	taskFound, err := service.taskRepository.GetById(taskId)

	if err != nil || taskFound == nil {
		return nil, errors.New("task not found")
	}

	err = service.isParentTaskValid(payload.ParentTaskId, taskFound.Project.ID.String())

	if err != nil {
		return nil, err
	}

	err = service.isAssigneeValid(payload.AssigneeId, taskFound.Project.ID.String())

	if err != nil {
		return nil, err
	}

	task, err := service.taskRepository.UpdateTask(taskFound.ID.String(), payload)

	if err != nil {
		return nil, err
	}

	return task, nil
}

func (service *TaskService) PatchUpdateTask(taskId string, payload *PatchUpdateTaskValidation, userId string) error {
	taskFound, err := service.taskRepository.GetById(taskId)

	if err != nil || taskFound == nil {
		return errors.New("task not found")
	}

	err = service.isParentTaskValid(payload.ParentTaskId, taskFound.Project.ID.String())

	if err != nil {
		return err
	}

	err = service.isAssigneeValid(payload.AssigneeId, taskFound.Project.ID.String())

	if err != nil {
		return err
	}

	err = service.taskRepository.PatchUpdateTask(taskFound.ID.String(), payload)

	if err != nil {
		return err
	}

	return nil
}

func (service *TaskService) DeleteTask(taskId string, userId string) error {
	task, err := service.taskRepository.GetById(taskId)

	if err != nil || task == nil {
		return errors.New("task not found")
	}
	return nil
}

func (service *TaskService) isParentTaskValid(parentTaskId uuid.NullUUID, projectId string) error {
	if !parentTaskId.Valid {
		return nil
	}

	if reflect.ValueOf(parentTaskId).IsZero() {
		return nil
	}

	parentTask, err := service.taskRepository.GetById(parentTaskId.UUID.String())

	if err != nil {
		return errors.New("parent task not found")
	}

	if parentTask != nil || parentTask.Project.ID.String() != projectId {
		return errors.New("parent task is not belong to this project")
	}

	return nil
}

func (service *TaskService) isAssigneeValid(assigneeId string, projectId string) error {
	if assigneeId == "" {
		return nil
	}

	projectMembers, err := service.projectService.GetListMember(projectId)

	if err != nil {
		return err
	}

	if !project.IsMember(projectMembers, assigneeId) {
		return errors.New("trying to assign task to someone who is not member of this project")
	}

	return nil
}

func (service *TaskService) GetListTask(userId string, query GetListTaskValidation) (tasks []*entities.Task, pagination *common.PaginationResponse, err error) {
	projectMembers, err := service.projectService.GetListMember(query.ProjectId)

	if err != nil || len(*projectMembers) == 0 {
		return nil, nil, errors.New("project not found")
	}

	if !project.IsMember(projectMembers, userId) {
		return nil, nil, errors.New("you are not member of this project")
	}

	tasks, total, err := service.taskRepository.GetListTask(query)

	if err != nil {
		return
	}

	pagination = utils.GetPaginationResponse(total, &query.Pagination)

	return
}

func (service *TaskService) UpdateOrderTasks(projectId string, sectionId string, userId string, tasks []string) error {
	projectMembers, err := service.projectService.GetListMember(projectId)

	if err != nil || len(*projectMembers) == 0 {
		return errors.New("project not found")
	}

	if !project.IsMember(projectMembers, userId) {
		return errors.New("you are not member of this project")
	}

	return service.taskRepository.UpdateOrderTasks(projectId, sectionId, tasks)

}
