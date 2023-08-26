package entities

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	BaseEntity
	Name        string    `gorm:"not null" json:"name"`
	CreatedById uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	CreatedBy   User      `json:"created_by"`
	Users       []User    `gorm:"many2many:project_users;" json:"users"`
	Tasks       []Task    `gorm:"foreignKey:ProjectId" json:"tasks"`
}

type ProjectUsers struct {
	UserId    uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	ProjectId uuid.UUID `gorm:"type:uuid;not null" json:"project_id"`
	Role      string    `gorm:"not null;default:member" json:"role"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
