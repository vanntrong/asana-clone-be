package entities

import (
	"github.com/google/uuid"
)

type Comment struct {
	BaseEntity
	Content  string    `gorm:"not null" json:"content"`
	AuthorId uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	TaskId   uuid.UUID `gorm:"type:uuid;not null" json:"-"`
	Author   *User     `json:"author"`
	Task     *Task     `json:"task"`
}
