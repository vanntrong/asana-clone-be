package comments

import (
	"fmt"

	"github.com/vanntrong/asana-clone-be/common"
	"github.com/vanntrong/asana-clone-be/entities"
	"github.com/vanntrong/asana-clone-be/project"
	"github.com/vanntrong/asana-clone-be/task"
	"github.com/vanntrong/asana-clone-be/utils"
)

type ICommentsService interface {
	Create(userId string, payload CreateCommentValidation) (*entities.Comment, error)
	GetListByTaskId(userId string, payload GetListByTaskIdValidation) (*[]entities.Comment, *common.PaginationResponse, error)
	Update(userId string, commentId string, payload UpdateCommentValidation) (*entities.Comment, error)
	Delete(userId string, commentId string) error
}

type CommentsService struct {
	repo           ICommentsRepository
	taskService    task.ITaskService
	projectService project.IProjectService
}

func NewCommentsService(repo ICommentsRepository, taskService task.ITaskService, projectService project.IProjectService) *CommentsService {
	return &CommentsService{repo: repo, taskService: taskService, projectService: projectService}
}

func (service *CommentsService) Create(userId string, payload CreateCommentValidation) (comment *entities.Comment, err error) {
	task, err := service.taskService.GetById(payload.TaskId, userId)

	if err != nil || task == nil {
		return nil, fmt.Errorf("Task not found")
	}

	comment, err = service.repo.Create(userId, payload)
	return
}

func (service *CommentsService) GetListByTaskId(userId string, payload GetListByTaskIdValidation) (comments *[]entities.Comment, pagination *common.PaginationResponse, err error) {
	task, err := service.taskService.GetById(payload.TaskId, userId)

	if err != nil || task == nil {
		return nil, nil, fmt.Errorf("Task not found")
	}

	projectMember, err := service.projectService.FindMember(task.ProjectId.String(), userId)

	if err != nil || projectMember == nil {
		return nil, nil, fmt.Errorf("You are not member of this project")
	}

	comments, total, err := service.repo.GetListByTaskId(payload)

	pagination = utils.GetPaginationResponse(total, &payload.Pagination)

	return
}

func (service *CommentsService) Update(userId string, commentId string, payload UpdateCommentValidation) (comment *entities.Comment, err error) {
	comment, err = service.repo.GetById(commentId)

	if err != nil || comment == nil {
		return nil, fmt.Errorf("Comment not found")
	}

	projectMember, err := service.projectService.FindMember(comment.Task.ProjectId.String(), userId)

	if err != nil || projectMember == nil {
		return nil, fmt.Errorf("You are not member of this project")
	}

	comment, err = service.repo.Update(commentId, payload)
	return
}

func (service *CommentsService) Delete(userId string, commentId string) (err error) {
	comment, err := service.repo.GetById(commentId)

	if err != nil || comment == nil {
		return fmt.Errorf("Comment not found")
	}

	projectMember, err := service.projectService.FindMember(comment.Task.ProjectId.String(), userId)

	if err != nil || projectMember == nil {
		return fmt.Errorf("You are not member of this project")
	}

	err = service.repo.Delete(commentId)
	return
}
