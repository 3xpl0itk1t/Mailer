package server

import (
	"csi_mailer/handlers"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func StartServer() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to the database
	handlers.ConnectToDB()
	defer handlers.DisconnectFromDB()

	// Start the server
	PORT := os.Getenv("PORT")
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}​\n",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, world!",
		})
	})
	app.Post("/signup", handlers.SignupHandler)
	app.Get("/login", handlers.LoginHandler)
	app.Post("/sendmail", handlers.MailHandler)
	// Run the server in a goroutine
	go func() {
		err := app.Listen(":" + PORT)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Wait for termination signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Server is quitting, disconnect from the database
	handlers.DisconnectFromDB()
}
