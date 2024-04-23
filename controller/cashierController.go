package controller

import (
	"github.com/gofiber/fiber/v2"
	"log"
	db "rest_api/config"
	"rest_api/models"
)

func CreateCashier(c *fiber.Ctx) error {
	var data map[string]string
	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("registeration error in post request %v", err)
	}

	if data["name"] == "" || data["passcode"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Name is required",
			"error":   map[string]interface{}{},
		})
	}

	cashier := models.Cashier{
		Name:     data["name"],
		Passcode: data["passcode"],
	}
	db.DB.Create(&cashier)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    cashier,
	})
}

func EditCashier(c *fiber.Ctx) error {

	return nil
}

func UpdateCashier(c *fiber.Ctx) error {

	return nil
}

func CashierList(c *fiber.Ctx) error {
	var cashier []models.Cashier
	//var count int64
	//limit, _ := strconv.Atoi(c.Query("limit"))
	//skip, _ := strconv.Atoi(c.Query("skip"))
	//db.DB.Select("*").Limit(limit).Offset(skip).Find(&cashier).Count(&count)
	db.DB.Select("*").Find(&cashier)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    cashier,
	})

}

func GetCashierDetails(c *fiber.Ctx) error {

	return nil
}

func DeleteCashier(c *fiber.Ctx) error {
	cashierId := c.Params("cashierId")
	var cashier models.Cashier
	db.DB.Where("id=?", cashierId).First(&cashier)

	if cashier.Id == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
			"error":   map[string]interface{}{},
		})
	}
	db.DB.Where("id = ?", cashierId).Delete(&cashier)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
	})
}
