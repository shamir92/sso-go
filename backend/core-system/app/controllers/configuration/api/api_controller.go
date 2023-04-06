package api

import (
	"reflect"
	"sso-go/app/models"
	"sso-go/platform/database"
	"strings"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func ValidateStruct(data models.API) []*models.ErrorResponse {
	var validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	var errors []*models.ErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element models.ErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.ActualTag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

func PostConfigurationAPI(c *fiber.Ctx) error {

	db := database.DBConn
	// var err error
	tempData := new(models.API)
	if err := c.BodyParser(tempData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := ValidateStruct(*tempData)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	if models.IsValidIdentifier(tempData.SigningAlgorithm) == false {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "unrecognize value of signing_algorithm",
		})
	}

	tempData.ID, _ = uuid.NewRandom()
	tempData.CreatedAt = time.Now()

	if err := db.Clauses().Create(&tempData).Error; err != nil {
		// error handling...
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg":  nil,
		"data": &tempData,
	})
}

func GetConfigurationAPI(c *fiber.Ctx) error {
	db := database.DBConn
	// var err error
	tx := []models.API{}
	if err := db.Find(&tx).Error; err != nil {
		// error handling...
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": nil,
		"data":    tx,
	})
}

func GetConfigurationAPIDetail(c *fiber.Ctx) error {
	db := database.DBConn
	uuidAPI := c.Params("uuid")
	tx := models.API{}
	if err := db.Where("ID = ?", uuidAPI).First(&tx).Error; err != nil {
		// error handling...
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": nil,
		"data":    tx,
	})
}
