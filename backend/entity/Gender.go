package entity

import "gorm.io/gorm"

type Gender struct {
	gorm.Model
	GenderName	string	`valid:"required~GenderName is required"`

	User	[]User	`gorm:"foreignKey:GenderID"`
}