package model

import (
	"time"
)

type GistModel struct {
	ID          string `gorm:"primaryKey"`
	OwnerID     string
	Title       string
	Description string
	Code        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
