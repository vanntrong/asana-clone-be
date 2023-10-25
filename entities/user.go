package entities

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	BaseEntity
	Name           string     `gorm:"not null;index" json:"name"`
	Email          string     `gorm:"unique;not null;index" json:"email" `
	Password       string     `gorm:"" json:"-"`
	Avatar         string     `gorm:"default:null" json:"avatar"`
	Provider       string     `gorm:"default:local" json:"provider"`
	ProjectCreated *[]Project `gorm:"foreignKey:CreatedById" json:"project_created,omitempty"`
	TaskAssigned   *[]Task    `gorm:"foreignKey:AssigneeId" json:"task_assigned,omitempty"`
	TaskCreated    *[]Task    `gorm:"foreignKey:CreatedById" json:"task_created,omitempty"`
	CommentCreated *[]Comment `gorm:"foreignKey:AuthorId" json:"comment_created,omitempty"`
	IsActive       bool       `gorm:"default:true" json:"is_active"`
}

type UserResponse struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email" `
	Avatar    string     `json:"avatar"`
	IsActive  bool       `json:"is_active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func (user *User) UserSerializer() *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Avatar:    user.Avatar,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}

func (user *User) HashPassword() *User {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	user.Password = string(hashedPassword)

	return user
}

func (user *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return false
	}

	return true
}
