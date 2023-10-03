package entities

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	BaseEntity
	Title        string        `gorm:"not null;index" json:"title"`
	Description  string        `gorm:"" json:"description"`
	StartDate    time.Time     `gorm:"not null;default:current_timestamp" json:"start_date"`
	DueDate      time.Time     `gorm:"not null" json:"due_date"`
	IsDone       bool          `gorm:"not null;default:false" json:"is_done"`
	Tags         string        `gorm:"" json:"tags"`
	Order        int64         `gorm:"index" json:"order"`
	AssigneeId   uuid.UUID     `gorm:"type:uuid;not null" json:"assignee_id"`
	ProjectId    uuid.UUID     `gorm:"type:uuid;not null" json:"project_id"`
	CreatedById  uuid.UUID     `gorm:"type:uuid;not null" json:"created_by_id"`
	ParentTaskId uuid.NullUUID `gorm:"type:uuid;default:null" json:"parent_task_id,dive,omitempty"`
	SectionId    uuid.UUID     `gorm:"type:uuid;default:null" json:"section_id,omitempty"`
	Assignee     *User         `json:"assignee,omitempty"`
	Project      *Project      `json:"project,omitempty"`
	CreatedBy    *User         `json:"created_by,dive,omitempty"`
	ParentTask   *Task         `json:"parent_task,omitempty"`
	Section      *Section      `json:"section,omitempty"`
	Comments     *[]Comment    `gorm:"foreignKey:TaskId" json:"comments,omitempty"`
}
