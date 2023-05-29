package routes

import (
	"sso-go/app/controllers/configuration/api"
	"sso-go/app/controllers/configuration/language"
	"sso-go/app/controllers/configuration/region"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	// route := a.Group("/api/v1")

	// Routes for GET method:
	// route.Get("/books", controllers.GetBooks)   // get list of all books
	// route.Get("/book/:id", controllers.GetBook) // get one book by ID

	// Routes for POST method:
	// route.Post("/user/sign/up", controllers.UserSignUp) // register a new user
	// route.Post("/user/sign/in", controllers.UserSignIn) // auth, return Access & Refresh tokens

	routeConfiguration := a.Group("configuration/v1")
	routeConfiguration.Post("/api", api.PostConfigurationAPI)
	routeConfiguration.Get("/api", api.GetConfigurationAPI)

	// ini harus ada refactor untuk :uuid ke :uuid_api
	routeConfiguration.Get("/api/:uuid", api.GetConfigurationAPIDetail)
	routeConfiguration.Put("/api/:uuid/setting", api.PostConfigurationAPISetting)
	routeConfiguration.Get("/api/:uuid/setting", api.GetConfigurationAPISetting)
	routeConfiguration.Post("/api/:uuid/permission", api.PostConfigurationAPIPermission)
	routeConfiguration.Get("/api/:uuid/permission", api.GetConfigurationAPIPermission)
	routeConfiguration.Get("/api/:uuid/permission/:uuid_api_permission", api.GetConfigurationAPIPermissionDetail)
	routeConfiguration.Delete("/api/:uuid/permission/:uuid_api_permission", api.DeleteConfigurationAPIPermissionDetail)
	routeConfiguration.Delete("/api/:uuid/permission/:uuid_api_permission/permanent", api.DeleteConfigurationAPIPermissionDetailPermanent)

	routeConfiguration.Get("/language", language.GetConfigurationLanguage)
	routeConfiguration.Delete("/language/:uuid_language", language.DeleteLanguage)
	routeConfiguration.Delete("/language/:uuid_language/permanent", language.DeleteLanguagePermanent)
	routeConfiguration.Get("/region", region.GetConfigurationRegion)
	routeConfiguration.Delete("/region/:uuid_region", region.DeleteRegion)
	routeConfiguration.Delete("/region/:uuid_region/permanent", region.DeleteRegionPermanent)

	routeSet := a.Group("initiator/v1")
	routeSet.Post("/language", language.SetInitDatabaseLanguage)
	routeSet.Post("/region", region.SetInitDatabaseRegion)

}
