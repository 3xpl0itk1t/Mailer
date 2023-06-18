package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var (
	AUTHENTICATION string = os.Getenv("AUTH_CODE")
	SECRET_KEY     string = os.Getenv("SECRET_KEY")
)
