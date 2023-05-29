package main

import (
	"fmt"
	"os"

	"sso-go/app/helpers"
	"sso-go/pkg/configs"
	"sso-go/pkg/middleware"
	"sso-go/pkg/routes"
	"sso-go/pkg/utils"
	"sso-go/platform/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "sso-go/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)
	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// initialization startup
	initDatabase()

	// Routes.
	// routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app) // Register a public routes for app.
	// routes.PrivateRoutes(app) // Register a private routes for app.
	// routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start server (with or without graceful shutdown).
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}

func initDatabase() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
		os.Getenv("DB_TIMEZONE"),
	)

	// database.DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })

	database.DBConn, err = gorm.Open(postgres.Open(dsn))
	sqlDb, err := database.DBConn.DB()

	sqlDb.SetMaxIdleConns(helpers.GetEnvInt("DB_MAX_IDLE_CONNS"))
	sqlDb.SetMaxOpenConns(helpers.GetEnvInt("DB_MAX_CONNECTIONS"))
	sqlDb.SetConnMaxLifetime(helpers.GetEnvTimeDuration("DB_MAX_LIFETIME_CONNECTIONS"))

	if err != nil {
		panic("failed to connect database")
	}

	// fmt.Println("Connection Opened to Database")
	// database.DBConn.AutoMigrate(
	// 	&models.API{},
	// 	&models.User{},
	// 	&models.Role{},
	// 	&models.APIPermission{},
	// 	&models.RolePermission{},
	// 	&models.UserRole{},
	// 	&models.Language{},
	// 	&models.Region{},
	// )
	// fmt.Println("Database Migrated")
}
