package main

import "github.com/gofiber/fiber/v2"

func main() {
	//1. fiber 인스턴스 생성
	app := fiber.New()

	//2. httphandler 생성
	// app.Get(A,B)
	// A = path string
	// B = handler Handler
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"sussess": true,
			"message": "Hello go Lang",
		})
	})
	//3. listen on port
	app.Listen(":3000")
}
