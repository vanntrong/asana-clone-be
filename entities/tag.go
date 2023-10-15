package entities

import "github.com/google/uuid"

type Tag struct {
	BaseEntity
	Name      string    `gorm:"not null;index" json:"name"`
	Color     string    `gorm:"not null;index" json:"color"`
	ProjectId uuid.UUID `gorm:"type:uuid;not null" json:"project_id"`
	Project   *Project  `json:"project,omitempty"`
}
