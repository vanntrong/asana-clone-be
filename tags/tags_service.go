package tags

import (
	"fmt"

	"github.com/vanntrong/asana-clone-be/entities"
	"github.com/vanntrong/asana-clone-be/project"
)

type ITagsService interface {
	FindTags(userId string, projectId string) (tags []*entities.Tag, err error)
	CreateTag(userId string, payload CreateTagValidation) (tag *entities.Tag, err error)
	UpdateTag(tagId string, userId string, payload UpdateTagValidation) (tag *entities.Tag, err error)
	AddTagToTask(userId string, payload AddTagToTaskValidation) (taskTag *entities.TaskTags, err error)
}

type TagsService struct {
	repo           ITagsRepository
	projectService project.IProjectService
}

func NewTagsService(repo ITagsRepository, projectService project.IProjectService) *TagsService {
	return &TagsService{repo, projectService}
}

func (s *TagsService) FindTags(userId string, projectId string) (tags []*entities.Tag, err error) {
	// Check if project exist and user is member of project
	projectUser, err := s.projectService.FindMember(projectId, userId)

	if err != nil || projectUser == nil {
		return nil, fmt.Errorf("project not found, or user is not member of project")
	}

	return s.repo.FindTags(projectId)
}

func (s *TagsService) CreateTag(userId string, payload CreateTagValidation) (tag *entities.Tag, err error) {
	// Check if project exist and user is member of project
	projectUser, err := s.projectService.FindMember(payload.ProjectId, userId)

	if err != nil || projectUser == nil {
		return nil, fmt.Errorf("project not found, or user is not member of project")
	}

	return s.repo.CreateTag(payload)
}

func (s *TagsService) UpdateTag(tagId string, userId string, payload UpdateTagValidation) (tag *entities.Tag, err error) {
	// check if tag exist
	tag, err = s.repo.FindTag(tagId)

	if err != nil || tag == nil {
		return nil, fmt.Errorf("tag not found")
	}

	// Check if project exist and user is member of project
	projectUser, err := s.projectService.FindMember(tag.ProjectId.String(), userId)

	if err != nil || projectUser == nil {
		return nil, fmt.Errorf("project not found, or user is not member of project")
	}

	return s.repo.UpdateTag(tagId, payload)
}

func (s *TagsService) AddTagToTask(userId string, payload AddTagToTaskValidation) (taskTag *entities.TaskTags, err error) {
	// check if tag exist
	tag, err := s.repo.FindTag(payload.TagId)

	if err != nil || tag == nil {
		return nil, fmt.Errorf("tag not found")
	}

	// Check if project exist and user is member of project
	projectUser, err := s.projectService.FindMember(tag.ProjectId.String(), userId)

	if err != nil || projectUser == nil {
		return nil, fmt.Errorf("project not found, or user is not member of project")
	}

	return s.repo.AddTagToTask(payload)
}
