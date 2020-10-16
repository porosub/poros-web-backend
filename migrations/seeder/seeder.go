package seeder

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/user"
	"github.com/divisi-developer-poros/poros-web-backend/models/user_type"
)

// Execute menjalankan seeder yang telah dibuat
func Execute() {
	var userTypes []user_type.User_Type
	if err := user_type.GetAll(&userTypes); err != nil {
		return
	}
	var users []user.User
	if err := user.GetAll(&users); err != nil {
		return
	}
	if len(userTypes) == 0 && len(users) == 0 {
		userTyp := user_type.User_Type{
			Id:   1,
			Name: "admin",
		}
		usr := user.User{
			Id:           1,
			Full_name:    "Root Access",
			Username:     "root",
			Password:     "12345",
			User_type_id: 1,
		}
		user_type.Create(&userTyp)
		user.Create(&usr)
	}

}
