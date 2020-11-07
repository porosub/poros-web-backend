package seeder

import (
	"github.com/divisi-developer-poros/poros-web-backend/config"
)

var (
	mysql      config.DBMySQL
	connection = mysql.MysqlConn()
)

// Execute menjalankan seeder yang telah dibuat
func Execute() {
	TagSeeder()
	PostTypeSeeder()
	UserTypeSeeder()
	UserSeeder()
}
