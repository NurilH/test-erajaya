package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model
	UsersID     uint
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
	Price       int    `json:"price" form:"price"`
	Quantity    int    `json:"quantity" form:"quantity"`
}

type ResProducts struct {
	ID          uint
	Name        string
	Description string
	Price       int
	Quantity    int
}
