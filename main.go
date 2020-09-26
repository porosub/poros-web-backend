package main

import (
	"github.com/divisi-developer-poros/poros-web-backend/config"
	"github.com/divisi-developer-poros/poros-web-backend/migrations"
	"github.com/divisi-developer-poros/poros-web-backend/routes"
	"log"
)

func main() {
	db, err := config.MysqlConn()
	if err != nil {
		log.Fatal(err)
	}
	migrations.Start(db)
	routes.Start()
}
