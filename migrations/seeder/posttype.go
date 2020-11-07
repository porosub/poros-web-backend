package seeder

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/posttype"
)

// PostTypeSeeder ... Post Type Seeder
func PostTypeSeeder() {
	var count int64
	if connection.Model(&posttype.PostType{}).Count(&count); count == 0 {
		connection.Create(&posttype.PostType{
			Name: "Berita",
		})
		connection.Create(&posttype.PostType{
			Name: "Tutorial",
		})
		connection.Create(&posttype.PostType{
			Name: "Acara",
		})
		connection.Create(&posttype.PostType{
			Name: "Artikel",
		})
		println("Post Type Seeder Executed Successfully")
	}
}
