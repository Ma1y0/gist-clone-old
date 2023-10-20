package model

import (
	"time"
)

type UserModel struct {
	ID        string `gorm:"primaryKey"`
	Username  string `gorm:"uniqueIndex"`
	Password  string
	Email     string      `gorm:"uniqueIndex"`
	Gists     []GistModel /* `gorm:"foreignKey:OwnerID"` */
	CreatedAt time.Time
	UpdatedAt time.Time
}
