package entities

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	BaseEntity
	IsDeleted    bool      `gorm:"default:false" json:"is_deleted"`
	DeletedAt    time.Time `gorm:"index" json:"deleted_at"`
	Title        string    `gorm:"not null;index" json:"title"`
	Description  string    `gorm:"not null" json:"description"`
	DueDate      time.Time `gorm:"not null" json:"due_date"`
	Status       string    `gorm:"not null" json:"status"`
	Tags         string    `gorm:"not null" json:"tags"`
	AssigneeId   uuid.UUID `gorm:"type:uuid;not null" json:"assignee_id"`
	ProjectId    uuid.UUID `gorm:"type:uuid;not null" json:"project_id"`
	CreatedById  uuid.UUID `gorm:"type:uuid;not null" json:"created_by_id"`
	ParentTaskId uuid.UUID `gorm:"type:uuid;default:null" json:"parent_task_id,omitempty"`
	Assignee     User      `json:"assignee,omitempty"`
	Project      Project   `json:"project,omitempty"`
	CreatedBy    User      `json:"created_by,dive,omitempty"`
	ParentTask   *Task     `json:"parent_task,omitempty"`
	Comments     []Comment `gorm:"foreignKey:TaskId" json:"comments,omitempty"`
}
