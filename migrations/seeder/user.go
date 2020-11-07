package seeder

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/user"
)

// UserSeeder ... User seeder
func UserSeeder() {
	var count int64
	if connection.Model(&user.User{}).Count(&count); count == 0 {
		usr := user.User{
			ID:         1,
			FullName:   "Root User",
			Username:   "root",
			Password:   "123456",
			UserTypeID: 1,
		}
		user.Create(&usr, nil)
		println("User Seeder Executed Successfully")
	}
}
