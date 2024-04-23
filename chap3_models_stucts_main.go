package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
	"rest_api/models"
)

var DB *gorm.DB

func main() {
	fmt.Println("db_connect...")
	godotenv.Load(".env")
	dbhost := os.Getenv("MYSQL_HOST")
	dbuser := os.Getenv("MYSQL_USER")
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbname)
	var db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("db connect fail")
	}
	DB = db
	fmt.Println("db connect success")

	AutoMigrate(db)

}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&models.Cashier{},
		&models.Category{},
		&models.Payment{},
		&models.PaymentType{},
		&models.Product{},
		&models.Discount{},
		&models.Order{},
	)
}
