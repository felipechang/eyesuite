package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gitlab.com/hardcake/eyesuite/controller"
	"gitlab.com/hardcake/eyesuite/middleware"
	"gitlab.com/hardcake/eyesuite/server"
	"gitlab.com/hardcake/eyesuite/service"
	"gitlab.com/hardcake/eyesuite/service/storage"
	"gitlab.com/hardcake/eyesuite/service/tesseract"
	"gitlab.com/hardcake/eyesuite/service/token"
	"log"
)

var tes tesseract.Tesseract
var tok token.Token
var sto storage.Storage

func init() {

	// load env module
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Get environment variables
	cts := service.GetEnv()
	log.Println("Environment variables loaded")

	// load services and service cors
	tes = tesseract.NewTesseract()
	log.Println("Tesseract module loaded")
	tok = token.NewToken(cts.AccessSecret, cts.RefreshSecret)
	log.Println("Token module loaded")

	sto = storage.NewStorage("db", "6379", "")
	log.Println("Redis module loaded")

	preLoadConfig()
	log.Println("Created dummy config")
	preProfile()
	log.Println("Created dummy profile")
	prePlugin()
	log.Println("Created dummy plugin")
	preLoadAdminUser()
	log.Println("Created dummy users")
}

func main() {

	s := service.NewService(sto, tes, tok)
	m := middleware.NewMiddleware(s)
	c := controller.NewController(s)

	app := fiber.New()

	// Serve static content
	app.Static("/", "./frontend/build")

	// Create a new route group '/api'
	api := app.Group("/api", c.ApiHome)

	// JWT management
	api.Post("/login", c.Login)
	api.Post("/logout", m.Auth, c.Logout)
	api.Post("/refresh", m.Auth, c.Refresh)

	// Post results to Netsuite
	api.Post("/postImage", m.Auth, c.PostImage)

	// Manage Users
	api.Get("/users", m.Auth, c.ReadUsers)
	api.Post("/users", m.Auth, m.Admin, c.UpsertUsers)

	// Manage Configuration
	api.Get("/config", m.Auth, c.ReadConfig)
	api.Post("/config", m.Auth, m.Admin, c.UpsertConfig)

	// Manage Profiles
	api.Get("/profiles", m.Auth, c.ReadProfiles)
	api.Post("/profiles", m.Auth, m.Admin, c.UpsertProfiles)

	// Manage Plugins
	api.Get("/plugins", m.Auth, c.ReadPlugins)
	api.Post("/plugins", m.Auth, m.Admin, c.UpsertPlugins)

	// Read image with Tesseract
	api.Post("/plugins/readText", m.Auth, c.ReadImageText)

	// All ready
	log.Println("Routes loaded")

	// Start!!
	server.Start(app)
}
