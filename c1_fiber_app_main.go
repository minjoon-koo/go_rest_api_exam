package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func main() {
	fmt.Println("db_connect...")
	godotenv.Load(".env")
	dbhost := os.Getenv("MYSQL_HOST")
	dbuser := os.Getenv("MYSQL_USER")
	dbpassword := os.Getenv("MYSQL_PASSWORD")
	dbname := os.Getenv("SALES_API_DB")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbname)
	var db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("db connect fail")
	}
	DB = db
	fmt.Println("db connect success")
}

/*
func main() {
	//1. fiber 인스턴스 생성
	app := fiber.New()

	//2. httphandler 생성
	// app.Get(A,B)
	// A = path string
	// B = handler Handler
	app.Get("/testApi", func(ctx fiber.Ctx) error {
		return ctx.Status(200).JSON(fiber.Map{
			"sussess": true,
			"message": "Hello go Lang",
		})
	})

	//3. listen on port
	app.Listen(":3000")
}
*/
