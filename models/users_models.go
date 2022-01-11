package models

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Token    string
	Products []Products
}

type ResLoginUser struct {
	ID       uint
	Email    string
	Token    string
	Password string
}
type ResCreateUser struct {
	ID       uint
	Email    string
	Password string
}
