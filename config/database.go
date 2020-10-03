package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBConfig ... Deklarasi atribut konfigurasi database
type DBConfig struct {
	Host     string
	Port     string
	User     string
	DBName   string
	Password string
}

type DBMySQL struct {
	// DB ... Referensi database yang digunakan
	DB gorm.DB
}

type DBMySQLInterface interface {
	BuildDBConfig() *DBConfig
	DbURL(dbConfig *DBConfig) string
	MysqlConn() (*gorm.DB, error)
}

// BuildDBConfig ... Inisialisasi konfigurasi pada database
func (dbsql *DBMySQL) BuildDBConfig() *DBConfig {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return &DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

// DbURL ... Mengambil url yang digunakan untuk driver
func (dbsql *DBMySQL) DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName, )
}


func (dbsql *DBMySQL) MysqlConn() *gorm.DB {
	result, err := gorm.Open(mysql.Open(dbsql.DbURL(dbsql.BuildDBConfig())), &gorm.Config{})
	if err != nil {
		return nil
	}
	return result
}