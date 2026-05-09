package repo

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitGrom() error {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		host := getEnv("MYSQL_HOST", "192.168.31.77")
		port := getEnv("MYSQL_PORT", "3306")
		user := getEnv("MYSQL_USER", "auto_call_phone")
		pass := getEnv("MYSQL_PASSWORD", "aiHakWpFm3HdBFyz")
		dbname := getEnv("MYSQL_DATABASE", "auto_call_phone")
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user, pass, host, port, dbname)
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func getEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
