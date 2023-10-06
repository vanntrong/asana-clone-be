package comments

import (
	"github.com/google/uuid"
	"github.com/vanntrong/asana-clone-be/entities"
	"github.com/vanntrong/asana-clone-be/utils"
	"gorm.io/gorm"
)

type ICommentsRepository interface {
	Create(userId string, payload CreateCommentValidation) (*entities.Comment, error)
	GetListByTaskId(payload GetListByTaskIdValidation) (*[]entities.Comment, int64, error)
	Update(commentId string, payload UpdateCommentValidation) (*entities.Comment, error)
	GetById(commentId string) (*entities.Comment, error)
	Delete(commentId string) error
}

type CommentsRepository struct {
	db *gorm.DB
}

func NewCommentsRepository(db *gorm.DB) *CommentsRepository {
	return &CommentsRepository{db: db}
}

func (repo *CommentsRepository) Create(userId string, payload CreateCommentValidation) (comment *entities.Comment, err error) {
	comment = &entities.Comment{
		Content:  payload.Content,
		AuthorId: uuid.MustParse(userId),
		TaskId:   uuid.MustParse(payload.TaskId),
	}

	err = repo.db.Create(comment).Error

	return
}

func (repo *CommentsRepository) GetListByTaskId(payload GetListByTaskIdValidation) (comments *[]entities.Comment, total int64, err error) {
	skip := utils.GetSkipValue(payload.Page, payload.Limit)

	comments = &[]entities.Comment{}
	err = repo.db.Model(&entities.Comment{}).
		Preload("Author").
		Where("task_id = ?", payload.TaskId).
		Where("deleted_at IS NULL").
		Limit(payload.Limit).
		Offset(skip).
		Count(&total).
		Find(comments).Error

	return
}

func (repo *CommentsRepository) Update(commentId string, payload UpdateCommentValidation) (comment *entities.Comment, err error) {
	err = repo.db.Model(&comment).
		Where("id = ?", commentId).
		Update("content", payload.Content).Error

	return
}

func (repo *CommentsRepository) GetById(commentId string) (*entities.Comment, error) {
	comment := &entities.Comment{}

	err := repo.db.Model(&entities.Comment{}).
		Preload("Author").
		Preload("Task").
		Where("id = ?", commentId).
		First(comment).Error

	return comment, err
}

func (repo *CommentsRepository) Delete(commentId string) error {
	deletedAt := utils.GetTimeNow()
	return repo.db.Model(&entities.Comment{}).Update("deleted_at", deletedAt.String()).Error
}
