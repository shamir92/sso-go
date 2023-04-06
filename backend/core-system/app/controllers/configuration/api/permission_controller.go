package api

import (
	"log"
	"sso-go/app/models"
	"sso-go/platform/database"

	"github.com/gofiber/fiber/v2"
)

func PostConfigurationPermission(c *fiber.Ctx) error {

	db := database.DBConn
	var err error
	log.Println("before validation")
	tempData := new(models.API)
	if err := c.BodyParser(tempData); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	log.Println(db)
	log.Println(err)
	log.Println("lulus validation")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   nil,
	})
}
