package user

import (
	"github.com/divisi-developer-poros/poros-web-backend/config"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/jinzhu/gorm"

	"github.com/divisi-developer-poros/poros-web-backend/util/Hash"
)

func GetAll(users *[]User) (err error) {
	if err := config.DB.Find(users).Error; err != nil {
		return err
	}
	return err
}

func Get(user *User, id int) (err error) {
	if err = config.DB.Where("id=?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func Create(user *User) (err error) {
	hashedPassword := Hash.GetMD5Hash(user.Password)
	user.Password = hashedPassword

	if err := config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func Update(updatedUser *User, id int) (err error) {
	hashedPassword := Hash.GetMD5Hash(updatedUser.Password)

	var existedUser User
	config.DB.Where("id=?", id).First(&existedUser)

	existedUser.Image = updatedUser.Image
	existedUser.Username = updatedUser.Username
	existedUser.Password = hashedPassword
	existedUser.Full_name = updatedUser.Full_name
	existedUser.User_type_id = updatedUser.User_type_id

	if err := config.DB.Save(existedUser).Error; err != nil {
		return err
	}
	return nil
}


func Delete (user *User, id int) (err error) {
	if err := config.DB.Unscoped().Where("id=?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}