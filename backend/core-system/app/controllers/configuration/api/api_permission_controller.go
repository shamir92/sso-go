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

func validateStructAPIPermission(data models.APIPermissionRequest) []*models.ErrorResponse {
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

func PostConfigurationAPIPermission(c *fiber.Ctx) error {

	db := database.DBConn
	uuidAPI := c.Params("uuid")
	reqData := new(models.APIPermissionRequest)
	tempData := new(models.APIPermission)
	if err := c.BodyParser(reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := validateStructAPIPermission(*reqData)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	api := models.API{}
	if err := db.Where("ID = ?", uuidAPI).First(&api).Error; err != nil {
		// error handling...
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	tempData.ID, _ = uuid.NewRandom()
	tempData.CreatedAt = time.Now()
	tempData.APIID = api.ID
	tempData.Scope = reqData.Scope
	tempData.Description = reqData.Description

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

func GetConfigurationAPIPermission(c *fiber.Ctx) error {
	db := database.DBConn
	// var err error
	uuidAPI := c.Params("uuid")
	tx := []models.APIPermission{}
	if err := db.Where("api_id = ?", uuidAPI).Find(&tx).Error; err != nil {
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
func GetConfigurationAPIPermissionDetail(c *fiber.Ctx) error {
	db := database.DBConn
	uuidAPI := c.Params("uuid")
	uuidAPIPermission := c.Params("uuid_api_permission")
	tx := models.APIPermission{}
	if err := db.Where("api_id = ?", uuidAPI).Where("id = ?", uuidAPIPermission).First(&tx).Error; err != nil {
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

func DeleteConfigurationAPIPermissionDetail(c *fiber.Ctx) error {
	db := database.DBConn
	uuidAPI := uuid.Must(uuid.Parse(c.Params("uuid")))
	uuidAPIPermission := uuid.Must(uuid.Parse(c.Params("uuid_api_permission")))
	tx := models.APIPermission{}
	tx.ID = uuidAPIPermission
	tx.APIID = uuidAPI
	if err := db.Where("api_id = ?", uuidAPI).Where("id = ?", uuidAPIPermission).First(&tx).Error; err != nil {
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

func DeleteConfigurationAPIPermissionDetailPermanent(c *fiber.Ctx) error {
	db := database.DBConn
	uuidAPI := uuid.Must(uuid.Parse(c.Params("uuid")))
	uuidAPIPermission := uuid.Must(uuid.Parse(c.Params("uuid_api_permission")))
	tx := models.APIPermission{}
	tx.ID = uuidAPIPermission
	tx.APIID = uuidAPI

	if err := db.Unscoped().Where("api_id = ?", uuidAPI).Where("id = ?", uuidAPIPermission).First(&tx).Error; err != nil {
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
