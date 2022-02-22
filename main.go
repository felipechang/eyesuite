package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/redirect/v2"
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

	// load services and service cors
	tes = tesseract.NewTesseract()
	tok = token.NewToken(cts.AccessSecret, cts.RefreshSecret)

	sto = storage.NewStorage("db", "6379", "")
	//sto = storage.NewStorage("localhost", "6379", "")

	preLoadConfig()
	preProfile()
	prePlugin()
	preLoadAdminUser()
}

func main() {

	s := service.NewService(sto, tes, tok)
	m := middleware.NewMiddleware(s)
	c := controller.NewController(s)

	app := fiber.New()
	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/help":                "/",
			"/settings/connection": "/",
			"/settings/plugins":    "/",
			"/settings/profiles":   "/",
			"/settings/users":      "/",
		},
		StatusCode: 301,
	}))

	// Serve static content
	app.Static("/", "./frontend/build")

	// JWT management
	app.Post("/api/login", c.Login)
	app.Post("/api/logout", m.Auth, c.Logout)
	app.Post("/api/refresh", m.Auth, c.Refresh)

	// Post results to Netsuite
	app.Post("/api/postImage", m.Auth, c.PostImage)

	// Manage Users
	app.Get("/api/users", m.Auth, c.ReadUsers)
	app.Post("/api/users", m.Auth, m.Admin, c.UpsertUsers)

	// Manage Configuration
	app.Get("/api/config", m.Auth, c.ReadConfig)
	app.Post("/api/config", m.Auth, m.Admin, c.UpsertConfig)

	// Manage Profiles
	app.Get("/api/profiles", m.Auth, c.ReadProfiles)
	app.Post("/api/profiles", m.Auth, m.Admin, c.UpsertProfiles)

	// Manage Plugins
	app.Get("/api/plugins", m.Auth, c.ReadPlugins)
	app.Post("/api/plugins", m.Auth, m.Admin, c.UpsertPlugins)

	// Read image with Tesseract
	app.Post("/api/plugins/readText", m.Auth, c.ReadImageText)

	// Api root
	app.Get("/api", c.ApiHome)

	// Start!!
	server.Start(app)
}
