package user_type

import (
	"github.com/divisi-developer-poros/poros-web-backend/config"
)

var (
	mysql      config.DBMySQL
	connection = mysql.MysqlConn()
)

func GetAll(userType *[]User_Type) (err error) {
	if err := connection.Find(userType).Error; err != nil {
		return err
	}
	return err
}

func Get(userType *User_Type, id int) (err error) {
	if err = connection.Where("id=?", id).First(userType).Error; err != nil {
		return err
	}
	return nil
}

func Create(userType *User_Type) (err error) {
	if err := connection.Create(userType).Error; err != nil {
		return err
	}
	return nil
}

func Update(updatedUserType *User_Type, id int) (err error) {
	var existedUserType User_Type
	connection.Where("id=?", id).First(&existedUserType)
	existedUserType.Name = updatedUserType.Name

	if err := connection.Save(existedUserType).Error; err != nil {
		return err
	}
	return nil
}


func Delete (userType *User_Type, id int) (err error) {
	if err := connection.Unscoped().Where("id=?", id).Delete(&userType).Error; err != nil {
		return err
	}
	return nil
}