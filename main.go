package main

import (
	"github.com/gin-gonic/gin"
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

	// get environment variables
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
	preProfile()
	prePlugin()
	preLoadAdminUser()
	log.Println("Created dummy records")
}

func main() {

	s := service.NewService(sto, tes, tok)
	m := middleware.NewMiddleware(s)
	c := controller.NewController(s)

	gin.SetMode(gin.ReleaseMode)

	// create router
	router := gin.Default()

	// Serve static content
	router.Static("/", "./frontend/build")

	// JWT management
	router.POST("api/login", c.Login)

	router.POST("api/logout", m.Auth, c.Logout)
	router.POST("api/refresh", m.Auth, c.Refresh)

	// Post results to Netsuite
	router.POST("api/postImage", m.Auth, c.PostImage)

	// Manage Users
	router.GET("api/users", m.Auth, c.ReadUsers)
	router.POST("api/users", m.Auth, m.Admin, c.UpsertUsers)

	// Manage Configuration
	router.GET("api/config", m.Auth, c.ReadConfig)
	router.POST("api/config", m.Auth, m.Admin, c.UpsertConfig)

	// Manage Profiles
	router.GET("api/profiles", m.Auth, c.ReadProfiles)
	router.POST("api/profiles", m.Auth, m.Admin, c.UpsertProfiles)

	// Manage Plugins
	router.GET("api/plugins", m.Auth, c.ReadPlugins)
	router.POST("api/plugins", m.Auth, m.Admin, c.UpsertPlugins)

	// Read image with Tesseract
	router.POST("api/plugins/readText", m.Auth, c.ReadImageText)

	// API Home
	router.GET("api/", c.ApiHome)

	// All ready
	log.Println("Routes loaded")

	// Start!!
	server.Start(router)
}
