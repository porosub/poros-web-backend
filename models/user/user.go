package user

import (
	"mime/multipart"

	"github.com/divisi-developer-poros/poros-web-backend/config"
	"github.com/divisi-developer-poros/poros-web-backend/utils/hash"
	"github.com/divisi-developer-poros/poros-web-backend/utils/storage"
)

var (
	mysql      config.DBMySQL
	connection = mysql.MysqlConn()
)

// GetAll ... Get all users from DB
func GetAll(users *[]User) (err error) {
	if err := connection.Preload("UserType").Find(users).Error; err != nil {
		return err
	}
	return err
}

// Get ... Get single user from DB
func Get(user *User, id int) (err error) {
	if err = connection.Where("id=?", id).Preload("UserType").First(user).Error; err != nil {
		return err
	}
	return nil
}

// GetByUsername ... Get single user from DB using his/her username
func GetByUsername(user *User, username string) (err error) {
	if err = connection.Where("username = ?", username).Preload("UserType").First(user).Error; err != nil {
		return err
	}
	return nil
}

// Create ... create new user to DB
func Create(user *User, imageBlob *multipart.FileHeader) (err error) {
	imageName, err := storeImage(imageBlob)
	if err != nil {
		return err
	}

	user.Password = hash.GetSha1Hash(user.Password)
	user.Image = imageName
	if err := connection.Create(user).Error; err != nil {
		storage.RemoveFile(config.AssetUsersImages + user.Image)
		return err
	}
	return nil
}

// Update ... update single user in DB
func Update(user *User, id int, imageBlob *multipart.FileHeader) (err error) {
	var oldUser User
	if err := Get(&oldUser, id); err != nil {
		return err
	}

	imageName, err := storeImage(imageBlob)
	if err != nil {
		return err
	}

	user.Password = hash.GetSha1Hash(user.Password)
	user.Image = imageName
	oldImage := oldUser.Image
	if err := connection.Model(&oldUser).Updates(&user).Error; err != nil {
		storage.RemoveFile(config.AssetUsersImages + user.Image)
		return err
	}
	if oldUser.Image != oldImage {
		storage.RemoveFile(config.AssetUsersImages + oldImage)
	}
	*user = oldUser
	return nil
}

// Delete ... Delete single user from DB
func Delete(id int) (err error) {
	var user User
	if err := Get(&user, id); err != nil {
		return err
	}

	if err := connection.Delete(&user).Error; err != nil {
		return err
	}
	storage.RemoveFile(config.AssetUsersImages + user.Image)
	return nil
}

// SignIn ... Sign in user with username and password
func SignIn(user *User, username string, password string) (err error) {
	hashedPassword := hash.GetSha1Hash(password)
	if err = connection.Where("username = ? AND password = ?", username, hashedPassword).Preload("UserType").First(&user).Error; err != nil {
		return err
	}
	return nil
}

func storeImage(imageBlob *multipart.FileHeader) (string, error) {
	if imageBlob == nil {
		return "", nil
	}

	filename, err := storage.StoreFileBlob(imageBlob, config.AssetUsersImages)
	if err != nil {
		return "", err
	}
	return filename, nil
}
