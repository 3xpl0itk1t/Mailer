package server

import (
	"csi_mailer/handlers"
	"csi_mailer/middlewares"
	"log"
	"os"
	"os/signal"
	"syscall"

	"csi_mailer/config"

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

	// Note: This is just an example, please use a secure secret key
	jwt := middlewares.NewAuthMiddleware(config.SECRET_KEY)
	// Logger for logging requests
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, world!",
		})
	})
	app.Post("/signup", handlers.SignupHandler)
	app.Get("/login", handlers.LoginHandler)
	app.Post("/sendmail", jwt, handlers.MailHandler)
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
