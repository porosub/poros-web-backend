package user

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/user_type"
	_ "github.com/go-playground/validator/v10"
)

type User struct {
	Id           	uint					`gorm:"primaryKey" json:"id" binding:"required"`
	Image        	string					`gorm:"not null" json:"image" binding:"required"`
	Username     	string					`gorm:"not null" json:"username" binding:"required"`
	Password     	string					`gorm:"not null" json:"password" binding:"required" validate:"min=8"`
	Full_name    	string					`gorm:"not null" json:"full_name" binding:"required"`
	User_Type		user_type.User_Type
}

func (b *User) TableName() string {
	return "users"
}