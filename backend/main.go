package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gitlab.com/hardcake/eyesuite/controller"
	"gitlab.com/hardcake/eyesuite/server"
	"gitlab.com/hardcake/eyesuite/service"
	"gitlab.com/hardcake/eyesuite/service/storage"
	"gitlab.com/hardcake/eyesuite/service/tesseract"
	"gitlab.com/hardcake/eyesuite/service/token"
	"log"
	"time"
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
	//sto = storage.NewStorage("localhost", "6379", "")
	log.Println("Redis module loaded")

	preLoadConfig()
	preProfile()
	prePlugin()
	preLoadAdminUser()
	log.Println("Created dummy records")
}

func main() {

	s := service.NewService(sto, tes, tok)
	//m := middleware.NewMiddleware(s)
	c := controller.NewController(s)

	gin.SetMode(gin.ReleaseMode)

	// create router
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			//TODO
			"*",
			"http://client:8080",
			"https://client:8080",
		},
		AllowMethods:     []string{"OPTIONS", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// JWT management
	router.POST("/login", c.Login)

	router.POST("/logout" /*m.Auth,*/, c.Logout)
	router.POST("/refresh" /*m.Auth,*/, c.Refresh)

	// Post results to Netsuite
	router.POST("/postImage" /*m.Auth,*/, c.PostImage)

	// Manage Users
	router.GET("/users" /*m.Auth, m.Admin,*/, c.ReadUsers)
	router.POST("/users" /*m.Auth, m.Admin,*/, c.UpsertUsers)

	// Manage Configuration
	router.GET("/config" /*m.Auth, m.Admin,*/, c.ReadConfig)
	router.POST("/config" /*m.Auth, m.Admin,*/, c.UpsertConfig)

	// Manage Profiles
	router.GET("/profiles" /*m.Auth, m.Admin,*/, c.ReadProfiles)
	router.POST("/profiles" /*m.Auth, m.Admin,*/, c.UpsertProfiles)

	// Manage Plugins
	router.GET("/plugins" /*m.Auth, m.Admin,*/, c.ReadPlugins)
	router.POST("/plugins" /*m.Auth, m.Admin,*/, c.UpsertPlugins)

	// Read image with Tesseract
	router.POST("/plugins/readText" /*m.Auth,*/, c.ReadImageText)

	// API Home
	router.GET("/", c.ApiHome)

	log.Println("Routes loaded")

	// Start!!
	server.Start(router)
}
