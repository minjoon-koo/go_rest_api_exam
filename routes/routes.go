package routes

import (
	"github.com/gofiber/fiber/v2"
	"rest_api/controller"
)

func Setup(app *fiber.App) {
	app.Post("/cashiers/:cashierId/login", controller.Login)
	app.Get("/cashiers/:cashierId/logout", controller.Logout)
	app.Post("/cashiers/:cashierId/passcode", controller.Passcode)

	//cashier routes
	app.Post("/cashiers", controller.CreateCashier)
	app.Get("/cashier", controller.CashierList)
	app.Get("/cashier/:cashierID", controller.GetCashierDetails)
	app.Delete("/cashiers/:cashierId", controller.DeleteCashier)
	app.Put("/cashiers/:cashierId", controller.UpdateCashier)
}
