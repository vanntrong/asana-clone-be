package user

import (
	"github.com/vanntrong/asana-clone-be/entities"
	"github.com/vanntrong/asana-clone-be/utils"
	"gorm.io/gorm"
)

type IUserRepository interface {
	FindByEmail(email string) (user *entities.User, err error)
	FindById(id string) (user *entities.User, err error)
	CreateUser(email string, password string, name string) (user *entities.User, err error)
	GetList(query GetListUserQuery) ([]*entities.User, int64, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) FindByEmail(email string) (user *entities.User, err error) {
	user = &entities.User{}
	result := repo.db.Where("email = ?", email).First(user)
	err = result.Error

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return user, err

}

func (repo *UserRepository) CreateUser(email string, password string, name string) (user *entities.User, err error) {
	user = &entities.User{
		Email:    email,
		Password: password,
		Name:     name,
	}
	err = repo.db.Create(&user).Error

	user.HashPassword()
	return user, err
}

func (repo *UserRepository) FindById(id string) (user *entities.User, err error) {
	user = &entities.User{}
	result := repo.db.Where("id = ?", id).First(user)
	err = result.Error

	if result.RowsAffected == 0 {
		return nil, nil
	}

	return user, err
}

func (repo *UserRepository) GetList(query GetListUserQuery) (users []*entities.User, total int64, err error) {
	skip := utils.GetSkipValue(query.Page, query.Limit)

	err = repo.db.Model(&users).
		Where(
			repo.db.Model(&entities.User{}).Where("name ILIKE ?", "%"+query.Keyword+"%").
				Or("email ILIKE ?", "%"+query.Keyword+"%"),
		).
		Where("is_deleted = ?", false).
		Limit(query.Limit).
		Offset(skip).
		Find(&users).
		Count(&total).Error

	return
}
