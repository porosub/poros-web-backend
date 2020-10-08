package main

import (
	"github.com/divisi-developer-poros/poros-web-backend/config"
	"github.com/divisi-developer-poros/poros-web-backend/migrations"
	"github.com/divisi-developer-poros/poros-web-backend/routes"
)

var mysql config.DBMySQL

func main() {
	db:= mysql.MysqlConn()
	migrations.Start(db)
	routes.Start()
}
