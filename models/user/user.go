package user

import (
	"github.com/divisi-developer-poros/poros-web-backend/config"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/jinzhu/gorm"
)

func GetAll(users *[]User) (err error) {
	if err := config.DB.Find(users).Error; err != nil {
		return err
	}
	return err
}

func Get(user *User, id string) (err error) {
	if err = config.DB.Where("id=?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}



