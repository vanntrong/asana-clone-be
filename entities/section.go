package entities

import (
	"github.com/google/uuid"
)

type Section struct {
	BaseEntity
	Name        string    `gorm:"not null;index" json:"name"`
	Description string    `gorm:"not null" json:"description"`
	ProjectId   uuid.UUID `gorm:"type:uuid;not null" json:"project_id"`
	Project     *Project  `json:"project,omitempty"`
	Tasks       *[]Task   `gorm:"foreignKey:SectionId" json:"tasks,omitempty"`
}
