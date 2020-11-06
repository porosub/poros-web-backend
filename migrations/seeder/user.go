package seeder

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/user"
)

func UserSeeder() {
	var count int64
	if connection.Model(&user.User{}).Count(&count); count == 0 {
		usr := user.User{
			Id:           1,
			Full_name:    "Root User",
			Username:     "root",
			Password:     "123456",
			User_type_id: 1,
		}
		user.Create(&usr)
		println("User Seeder Executed Successfully")
	}
}
