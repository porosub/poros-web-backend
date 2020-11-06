package seeder

import (
	"github.com/divisi-developer-poros/poros-web-backend/models/base"
	"github.com/divisi-developer-poros/poros-web-backend/models/tags"
)

func TagSeeder() {
	var count int64
	if connection.Model(&tags.Tag{}).Count(&count); count == 0 {
		connection.Create(&tags.Tag{
			base.Tag{
				Name: "Litbang",
			},
		})
		connection.Create(&tags.Tag{
			base.Tag{
				Name: "Internal",
			},
		})
		connection.Create(&tags.Tag{
			base.Tag{
				Name: "Humas",
			},
		})
		connection.Create(&tags.Tag{
			base.Tag{
				Name: "Developer",
			},
		})
		connection.Create(&tags.Tag{
			base.Tag{
				Name: "Operations",
			},
		})
		connection.Create(&tags.Tag{
			base.Tag{
				Name: "Security",
			},
		})
		println("Tags Seeder Executed Successfully")
	}
}
