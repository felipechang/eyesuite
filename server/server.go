package server

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func Start(app *fiber.App) {
	go func() {
		log.Println("Starting server")
		if err := app.Listen(":8080"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	//Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	if err := app.Shutdown(); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
