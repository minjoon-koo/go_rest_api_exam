package main

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	db "rest_api/config"
	routes "rest_api/routes"
)

func main() {
	fmt.Println("Connecting to DB...")
	db.Connection()

	app := fiber.New()
	app.Use(cors.New())

	routes.Setup(app)

	app.Listen(":30001")

}
