package models

import (
	"gorm.io/gorm"
)

type Blacklist struct {
	gorm.Model
	Token string `gorm:"uniqueIndex"`
}
