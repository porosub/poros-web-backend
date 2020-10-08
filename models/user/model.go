package user

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/user_type"
	_ "github.com/go-playground/validator/v10"
)

type User struct {
	Id           	uint					`gorm:"primaryKey" json:"id"`
	Image        	string					`gorm:"not null" json:"image" form:"image"`
	Username     	string					`gorm:"not null" json:"username" form:"username" binding:"required"`
	Password     	string					`gorm:"not null" json:"password" form:"password" binding:"required"`
	Full_name    	string					`gorm:"not null" json:"full_name" form:"full_name" binding:"required"`
	User_type_id	int						`form:"user_type_id" binding:"required"`
	User_Type		user_type.User_Type		`gorm:"foreignKey:User_type_id"`
}

func (b *User) TableName() string {
	return "users"
}