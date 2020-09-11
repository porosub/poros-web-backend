package migrations

import (
	"github.com/divisi-developer-poros/poros-web-backend/migrations/seeder"
	"github.com/jinzhu/gorm"
)

// Start melakukan migrasi ke database
func Start(db *gorm.DB) {
	// Simpan migrasi modelmu dibawah
	db.AutoMigrate()

	// Normal query juga bisa disimpan disini

	// Mengeksekusi seeder
	seeder.Execute()
}
