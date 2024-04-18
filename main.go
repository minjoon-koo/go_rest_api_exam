package main

import (
	"fmt"
	"github.com/joho/godotenv"
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

	connection :=
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
