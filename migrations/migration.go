package migrations

import (
	"github.com/divisi-developer-poros/poros-web-backend/migrations/seeder"
	"github.com/jinzhu/gorm"

	"github.com/divisi-developer-poros/poros-web-backend/models/user"

)

// Start melakukan migrasi ke database
func Start(db *gorm.DB) {
	// Simpan migrasi modelmu dibawah
	db.AutoMigrate(&user.User{})

	// Normal query juga bisa disimpan disini

	// Mengeksekusi seeder
	seeder.Execute()
}
