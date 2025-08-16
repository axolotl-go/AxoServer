package model

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
