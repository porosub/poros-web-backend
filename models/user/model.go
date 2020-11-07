package user

import (
	"github.com/divisi-developer-poros/poros-web-backend/config"
	"github.com/divisi-developer-poros/poros-web-backend/models/usertype"
	"github.com/divisi-developer-poros/poros-web-backend/utils/host"
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
	UserType   usertype.UserType `json:"user_type" binding:"-" `
}

// LocalizedField ... Localized all field to corresponded host
func (u *User) LocalizedField() {
	if u.Image != "" {
		u.Image = host.GetURL() + config.AssetUsersImages + u.Image
	}
}

// TableName ... User Table Name in DB
func (u *User) TableName() string {
	return "users"
}
