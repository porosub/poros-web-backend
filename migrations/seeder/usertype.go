package seeder

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/user_type"
)

func UserTypeSeeder() {
	var count int64
	if connection.Model(&user_type.User_Type{}).Count(&count); count == 0 {
		connection.Create(&user_type.User_Type{
			Name: "Root",
		})
		connection.Create(&user_type.User_Type{
			Name: "Admin",
		})
		connection.Create(&user_type.User_Type{
			Name: "Member",
		})
		println("User Type Seeder Executed Successfully")
	}
}
