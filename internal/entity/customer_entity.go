package entity

import (
	"github.com/google/uuid"
)

type Customer struct {
	// gorm.Model
	ID       uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid()"`
	ParentId uuid.NullUUID `gorm:"type:uuid;"`
	Name     string        `json:"name" gorm:"index;not null;size:5000"`
	Nomor    string        `json:"nomor" gorm:"index;not null;type:char(50)"`
}
