package user

import (
	_ "github.com/go-playground/validator/v10"
)

type User struct {
	Id           uint		`gorm:"primaryKey" json:"id"`
	Image        string		`gorm:"not null" json:"image" validate:"required"`
	Username     string		`gorm:"not null" json:"username" validate:"required"`
	Password     string		`gorm:"not null" json:"password" validate:"required,min=8"`
	Full_name    string		`gorm:"not null" json:"full_name" validate:"required"`
	User_type_id uint		`gorm:"not null" gorm:"index" json:"user_type_id" validate:"required"`
}

func (b *User) TableName() string {
	return "users"
}