package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"Password"`

	//Relationships with different account types
	MainAccount    MainAccount
	SavingsAccount SavingsAccount
	CardAccount    CardAccount
}
