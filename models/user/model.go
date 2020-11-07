package user

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/usertype"
	_ "github.com/go-playground/validator/v10" // Validator
)

// User ... User model declaration
type User struct {
	ID         uint              `gorm:"primaryKey" json:"id"`
	Image      string            `gorm:"not null" json:"image" form:"image"`
	Username   string            `gorm:"not null;unique" json:"username" form:"username" binding:"required"`
	Password   string            `gorm:"not null" json:"password" form:"password" binding:"required"`
	FullName   string            `gorm:"not null" json:"full_name" form:"full_name" binding:"required"`
	UserTypeID int               `json:"user_type_id" form:"user_type_id" binding:"required"`
	UserType   usertype.UserType `gorm:"foreignKey:UserTypeID" binding:"-" json:"user_type"`
}

// TableName ... User Table Name in DB
func (b *User) TableName() string {
	return "users"
}
