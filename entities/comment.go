package entities

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	BaseEntity
	DeletedAt time.Time `gorm:"index" json:"deleted_at"`
	Content   string    `gorm:"not null" json:"content"`
	AuthorId  uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	TaskId    uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Author    *User     `json:"author"`
	Task      *Task     `json:"task"`
}
