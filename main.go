package main

import (
	"fmt"

	"github.com/divisi-developer-poros/poros-web-backend/config"
	"github.com/divisi-developer-poros/poros-web-backend/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/divisi-developer-poros/poros-web-backend/migrations"
)

func main() {
	var err error

	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status: ", err)
	}

	defer config.DB.Close()

	migrations.Start(config.DB)

	routes.Start()
}
