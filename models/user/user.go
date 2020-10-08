package user

import (
	"github.com/divisi-developer-poros/poros-web-backend/config"

	"github.com/divisi-developer-poros/poros-web-backend/utils/Hash"
)

var (
	mysql      config.DBMySQL
	connection = mysql.MysqlConn()
)

func GetAll(users *[]User) (err error) {
	if err := connection.Preload("User_Type").Find(users).Error; err != nil {
		return err
	}
	return err
}

func Get(user *User, id int) (err error) {
	if err = connection.Where("id=?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func Create(user *User) (err error) {
	hashedPassword := Hash.GetMD5Hash(user.Password)
	user.Password = hashedPassword

	if err := connection.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func Update(updatedUser *User, id int) (err error) {
	hashedPassword := Hash.GetMD5Hash(updatedUser.Password)

	var existedUser User
	connection.Where("id=?", id).First(&existedUser)

	existedUser.Image = updatedUser.Image
	existedUser.Username = updatedUser.Username
	existedUser.Password = hashedPassword
	existedUser.Full_name = updatedUser.Full_name
	existedUser.User_Type = updatedUser.User_Type

	if err := connection.Save(existedUser).Error; err != nil {
		return err
	}
	return nil
}


func Delete (user *User, id int) (err error) {
	if err := connection.Unscoped().Where("id=?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}