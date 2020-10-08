package user

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/user_type"
	_ "github.com/go-playground/validator/v10"
)

type User struct {
	Id           	uint					`gorm:"primaryKey" json:"id"`
	Image        	string					`gorm:"not null" json:"image"`
	Username     	string					`gorm:"not null" json:"username" binding:"required"`
	Password     	string					`gorm:"not null" json:"password" binding:"required"`
	Full_name    	string					`gorm:"not null" json:"full_name" binding:"required"`
	User_Type		user_type.User_Type		`json:"user_type"`
	User_type_id	int						`json:"user_type_id"`
}

func (b *User) TableName() string {
	return "users"
}