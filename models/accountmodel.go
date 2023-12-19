package models

import "gorm.io/gorm"

type MainAccount struct {
	gorm.Model
	UserID            uint    `gorm:"index;onDelete:CASCADE"`
	Balance           float64 `gorm:"not null;check:balance >= 0"`
	MainAccountNumber string
	Transactions      []MainAccountTransaction `gorm:"foreignKey:MainAccountID"`
}

type SavingsAccount struct {
	gorm.Model
	UserID               uint    `gorm:"index;onDelete:CASCADE"`
	Balance              float64 `gorm:"not null;check:balance >= 0"`
	SavingsAccountNumber string
	Transactions         []SavingsAccountTransaction `gorm:"foreignKey:SavingsAccountID"`
}

type CardAccount struct {
	gorm.Model
	UserID       uint                     `gorm:"index;onDelete:CASCADE"`
	Balance      float64                  `gorm:"not null;check:balance >= 0"`
	Transactions []CardAccountTransaction `gorm:"foreignKey:CardAccountID"`
}
