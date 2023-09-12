package models

import (
	"time"

	"github.com/google/uuid"
)

type ToDo struct {
	ID          uuid.UUID `json:"id,omitempty"  example:"2afbae48-4fa5-41f6-9a2f-82b9503d18ed" gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string    `json:"name" validate:"required" example:"ToDo14" gorm:"type:varchar(255);not null"`
	Description string    `json:"description" validate:"required" example:"My New ToDo" gorm:"type:varchar(255)"`
	Status      string    `json:"status" example:"New" gorm:"type:varchar(255)"`
	CreatedAt   time.Time `json:",omitempty" example:"2023-09-12T15:45:44.523792+03:00"`
	UpdatedAt   time.Time `json:",omitempty" example:"2023-09-12T15:45:44.523792+03:00"`
}
