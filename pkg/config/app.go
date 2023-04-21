package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:Shyam@04@(127.0.0.1:3306)/goserver?charset=utf8&parseTime=True"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	db = d

}
func GetDB() *gorm.DB {
	return db
}
