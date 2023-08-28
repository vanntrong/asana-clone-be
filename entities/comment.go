package entities

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	BaseEntity
	IsDeleted bool      `gorm:"default:false" json:"is_deleted"`
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
	Content   string    `gorm:"not null" json:"content"`
	AuthorId  uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Author    *User     `json:"author"`
	TaskId    uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Task      *Task     `json:"task"`
}
