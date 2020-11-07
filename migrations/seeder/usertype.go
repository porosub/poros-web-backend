package seeder

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/usertype"
)

// UserTypeSeeder ... User type seeder
func UserTypeSeeder() {
	var count int64
	if connection.Model(&usertype.UserType{}).Count(&count); count == 0 {
		connection.Create(&usertype.UserType{
			Name: "Root",
		})
		connection.Create(&usertype.UserType{
			Name: "Admin",
		})
		connection.Create(&usertype.UserType{
			Name: "Member",
		})
		println("User Type Seeder Executed Successfully")
	}
}
