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
	"strings"
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

	gin.SetMode(gin.ReleaseMode)

	// create router
	router := gin.Default()
	apiRouter := gin.Default()
	staticRouter := gin.Default()

	// Serve static content
	staticRouter.Static("/", "./frontend/build")

	// API routes group
	a := apiRouter.Group("/api")
	{
		// JWT management
		a.POST("login", c.Login)

		a.POST("logout", m.Auth, c.Logout)
		a.POST("refresh", m.Auth, c.Refresh)

		// Post results to Netsuite
		a.POST("postImage", m.Auth, c.PostImage)

		// Manage Users
		a.GET("users", m.Auth, c.ReadUsers)
		a.POST("users", m.Auth, m.Admin, c.UpsertUsers)

		// Manage Configuration
		a.GET("config", m.Auth, c.ReadConfig)
		a.POST("config", m.Auth, m.Admin, c.UpsertConfig)

		// Manage Profiles
		a.GET("profiles", m.Auth, c.ReadProfiles)
		a.POST("profiles", m.Auth, m.Admin, c.UpsertProfiles)

		// Manage Plugins
		a.GET("plugins", m.Auth, c.ReadPlugins)
		a.POST("plugins", m.Auth, m.Admin, c.UpsertPlugins)

		// Read image with Tesseract
		a.POST("plugins/readText", m.Auth, c.ReadImageText)

		// API Home
		a.GET("/", c.ApiHome)
	}

	router.Any("/*any", func(c *gin.Context) {
		path := c.Param("any")
		if strings.HasPrefix(path, "/api") {
			apiRouter.HandleContext(c)
		} else {
			staticRouter.HandleContext(c)
		}
		c.Abort()
	})

	// All ready
	log.Println("Routes loaded")

	// Start!!
	server.Start(router)
}
