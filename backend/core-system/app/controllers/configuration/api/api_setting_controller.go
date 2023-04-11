package api

import (
	"reflect"
	"sso-go/app/models"
	"sso-go/platform/database"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func validateStructAPISetting(data models.APISettingRequest) []*models.ErrorResponse {
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

func PostConfigurationAPISetting(c *fiber.Ctx) error {
	db := database.DBConn
	// var x bool
	// var err error
	uuidAPI := c.Params("uuid")
	reqData := new(models.APISettingRequest)
	if err := c.BodyParser(reqData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	errors := validateStructAPISetting(*reqData)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	tx := new(models.API)
	if err := db.Where("ID = ?", uuidAPI).First(&tx).Error; err != nil {
		// error handling...
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	tx.Name = (reqData.Name)
	tx.Identifier = (reqData.Identifier)
	tx.TokenExpiration = (reqData.TokenExpiration)
	tx.TokenExpirationForBrowser = (reqData.TokenExpirationForBrowser)
	tx.EnableRBAC, _ = strconv.ParseBool(reqData.EnableRBAC)
	tx.AddPermissionToAccessToken, _ = strconv.ParseBool(reqData.AddPermissionToAccessToken)
	tx.AllowSkippingUserConsent, _ = strconv.ParseBool(reqData.AllowSkippingUserConsent)
	tx.AllowOfflineAccess, _ = strconv.ParseBool(reqData.AllowOfflineAccess)

	if err := db.Clauses().Save(&tx).Error; err != nil {
		// error handling...
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
			"data":    nil,
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg":  nil,
		"data": &tx,
	})
}

func GetConfigurationAPISetting(c *fiber.Ctx) error {
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
