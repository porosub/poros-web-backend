package usertype

import (
	"github.com/divisi-developer-poros/poros-web-backend/config"
)

var (
	mysql      config.DBMySQL
	connection = mysql.MysqlConn()
)

// GetAll ... Get All User Types from DB
func GetAll(userType *[]UserType) (err error) {
	if err := connection.Find(userType).Error; err != nil {
		return err
	}
	return err
}

// Get ... Get single user type from DB
func Get(userType *UserType, id int) (err error) {
	if err = connection.Where("id=?", id).First(userType).Error; err != nil {
		return err
	}
	return nil
}

// Create ... Create single user type to DB
func Create(userType *UserType) (err error) {
	if err := connection.Create(userType).Error; err != nil {
		return err
	}
	return nil
}

// Update ... Update single user type from DB
func Update(updatedUserType *UserType, id int) (err error) {
	var existedUserType UserType
	connection.Where("id=?", id).First(&existedUserType)
	existedUserType.Name = updatedUserType.Name

	if err := connection.Save(existedUserType).Error; err != nil {
		return err
	}
	return nil
}

// Delete ... Delete single user type from DB
func Delete(userType *UserType, id int) (err error) {
	if err := connection.Unscoped().Where("id=?", id).Delete(&userType).Error; err != nil {
		return err
	}
	return nil
}
