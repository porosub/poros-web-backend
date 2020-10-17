package migrations

import (
	"github.com/divisi-developer-poros/poros-web-backend/migrations/seeder"
	"github.com/divisi-developer-poros/poros-web-backend/models/token"
	"github.com/divisi-developer-poros/poros-web-backend/models/user"
	"github.com/divisi-developer-poros/poros-web-backend/models/user_type"

	"github.com/divisi-developer-poros/poros-web-backend/models/tags"
	"gorm.io/gorm"
)

// Start melakukan migrasi ke database
func Start(db *gorm.DB) {
	// Manually adding that fcking foreign key

	// Simpan migrasi modelmu dibawah
	db.AutoMigrate(&user.User{}, &user_type.User_Type{})
	db.AutoMigrate(&tags.Tag{})
	db.AutoMigrate(&token.AccessToken{})

	// Normal query juga bisa disimpan disini

	// Mengeksekusi seeder
	seeder.Execute()
}
