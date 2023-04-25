package language

import (
	"reflect"
	"sso-go/app/models"
	"sso-go/platform/database"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func validateStructLanguage(data models.LanguageRequest) []*models.ErrorResponse {
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

func GetConfigurationLanguage(c *fiber.Ctx) error {
	db := database.DBConn
	// var err error
	tx := []models.Language{}
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

func SetInitDatabase(c *fiber.Ctx) error {
	db := database.DBConn

	tempData := new(models.LanguageSetJson)
	if err := c.BodyParser(tempData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var languages []models.Language

	for _, arr := range *tempData {
		// perform an operation
		language := new(models.Language)
		language.Name = arr.Name
		language.DisplayName = arr.Name
		language.LanguageCode = arr.Code
		language.ID, _ = uuid.NewRandom()
		languages = append(languages, *language)
	}
	db.CreateInBatches(languages, 100)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": nil,
		"data":    tempData,
	})
}

func DeleteLanguage(c *fiber.Ctx) error {
	db := database.DBConn
	uuidLanguage := uuid.Must(uuid.Parse(c.Params("uuid_language")))
	tx := models.Language{}

	if err := db.Where("ID = ?", uuidLanguage).First(&tx).Error; err != nil {
		// error handling...
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := db.Delete(&tx).Error; err != nil {
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

func DeleteLanguagePermanent(c *fiber.Ctx) error {
	db := database.DBConn
	uuidLanguage := uuid.Must(uuid.Parse(c.Params("uuid_language")))
	tx := models.Language{}
	tx.ID = uuidLanguage

	if err := db.Unscoped().Where("id = ?", uuidLanguage).First(&tx).Error; err != nil {
		// error handling...
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := db.Unscoped().Delete(&tx).Error; err != nil {
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
