package entity

import (

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	StudentID	string	`valid:"required~StudentID is required, matches(^[BMD]\\d{7}$)~StudentID is invalid format"`
	Name	string	`valid:"required~Name is required, alpha~Name must be only letters"`
	Email	string	`valid:"required~Email is required, email~Invalid email"`
	PhoneNumber	string	`valid:"required~PhoneNumber is required, stringlength(10|10)"`
	GenderID	uint	`valid:"required~GenderID is required"`
	Gender		Gender	`gorm:"foreignKey:GenderID" valid:"-"`
}