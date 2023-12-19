package models

import (
	"time"

	"gorm.io/gorm"
)

type MainAccountTransaction struct {
	gorm.Model
	UserID        uint
	Amount        float64
	Description   string
	MainAccountID uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type SavingsAccountTransaction struct {
	gorm.Model
	UserID           uint
	Amount           float64
	Description      string
	SavingsAccountID uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type CardAccountTransaction struct {
	gorm.Model
	UserID        uint
	Amount        float64
	Description   string
	CardAccountID uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
