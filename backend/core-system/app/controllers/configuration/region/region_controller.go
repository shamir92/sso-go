package region

import (
	"reflect"
	"sso-go/app/models"
	"sso-go/platform/database"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func validateStructRegion(data models.LanguageRequest) []*models.ErrorResponse {
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

func GetConfigurationRegion(c *fiber.Ctx) error {
	db := database.DBConn
	// var err error
	tx := []models.Region{}
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

func SetInitDatabaseRegion(c *fiber.Ctx) error {
	db := database.DBConn

	tempData := new(models.RegionSetJson)
	if err := c.BodyParser(tempData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	var regions []models.Region

	for _, arr := range *tempData {
		// perform an operation
		region := new(models.Region)
		region.Name = arr.Name
		region.DisplayName = arr.Name
		region.RegionCode = arr.Code
		region.ID, _ = uuid.NewRandom()
		regions = append(regions, *region)
	}
	db.CreateInBatches(regions, 100)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": nil,
		"data":    tempData,
	})
}

func DeleteRegion(c *fiber.Ctx) error {
	db := database.DBConn
	uuidRegion := uuid.Must(uuid.Parse(c.Params("uuid_region")))
	tx := models.Region{}

	if err := db.Where("ID = ?", uuidRegion).First(&tx).Error; err != nil {
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

func DeleteRegionPermanent(c *fiber.Ctx) error {
	db := database.DBConn
	uuidRegion := uuid.Must(uuid.Parse(c.Params("uuid_region")))
	tx := models.Region{}
	tx.ID = uuidRegion

	if err := db.Unscoped().Where("id = ?", uuidRegion).First(&tx).Error; err != nil {
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
