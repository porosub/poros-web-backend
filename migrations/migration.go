package migrations

import (
	"github.com/divisi-developer-poros/poros-web-backend/migrations/seeder"
	"github.com/divisi-developer-poros/poros-web-backend/models/user"
	"github.com/divisi-developer-poros/poros-web-backend/models/tags"
	"gorm.io/gorm"
)

// Start melakukan migrasi ke database
func Start(db *gorm.DB) {
	// Simpan migrasi modelmu dibawah
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&tags.Tag{})

	// Normal query juga bisa disimpan disini

	// Mengeksekusi seeder
	seeder.Execute()
}
