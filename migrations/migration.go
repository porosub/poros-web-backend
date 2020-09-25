package migrations

import (
	"github.com/divisi-developer-poros/poros-web-backend/migrations/seeder"
	"github.com/jinzhu/gorm"

	"github.com/divisi-developer-poros/poros-web-backend/models/user"
	"github.com/divisi-developer-poros/poros-web-backend/models/user_type"

)

// Start melakukan migrasi ke database
func Start(db *gorm.DB) {
	// Simpan migrasi modelmu dibawah
	db.AutoMigrate(&user.User{}, &user_type.User_Type{})

	// Normal query juga bisa disimpan disini

	// Mengeksekusi seeder
	seeder.Execute()
}
