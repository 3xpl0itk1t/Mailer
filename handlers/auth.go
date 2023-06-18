package handlers

import (
	"context"
	"csi_mailer/config"
	"csi_mailer/models"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// SignupHandler is used to create a new user account
func SignupHandler(c *fiber.Ctx) error {
	var user models.SignupUser
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Check if user with the same email already exists
	existingUser := bson.M{"email": user.Email}
	count, err := collection.CountDocuments(context.Background(), existingUser)
	if err != nil {
		log.Fatal(err)
	}

	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "user with this email already exists",
		})
	}
	// Encrypt the password
	hashedAuth, err := bcrypt.GenerateFromPassword([]byte(user.Auth), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to encrypt password",
		})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to encrypt password",
		})
	}
	newUser := bson.M{
		"auth":       string(hashedAuth),
		"first_name": user.Firstname,
		"last_name":  user.Lastname,
		"email":      user.Email,
		"password":   string(hashedPassword),
	}
	if user.Auth != config.AUTHENTICATION {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Authentication failed",
		})
	}

	_, err = collection.InsertOne(context.Background(), newUser)
	if err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Signed Up successful!",
	})
}

func LoginHandler(c *fiber.Ctx) error {
	// Retrieve the user's login credentials
	var loginCredentials models.LoginUser
	if err := c.BodyParser(&loginCredentials); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Retrieve the stored hashed password from the database based on the user's email or username
	storedPassword, userID := getPasswordAndIDFromDatabase(loginCredentials.Email)

	// Compare the user's provided password with the stored hashed password
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(loginCredentials.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// Passwords match, generate a JWT token
	token := generateToken(userID.Hex())

	// Return the JWT token in the response
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful!",
		"token":   token,
	})
}

func getPasswordAndIDFromDatabase(email string) (string, primitive.ObjectID) {
	var result bson.M
	filter := bson.M{"email": email}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	hashedPassword := result["password"].(string)
	userID := result["_id"].(primitive.ObjectID)
	return hashedPassword, userID
}

func generateToken(userID string) string {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims (payload) for the token
	claims := token.Claims.(jwt.MapClaims)
	claims["_id"] = userID

	// Generate the JWT token with a secret key
	// Replace "your-secret-key" with your own secret key for signing the token
	tokenString, err := token.SignedString([]byte(config.SECRET_KEY))
	if err != nil {
		log.Fatal(err)
	}

	return tokenString
}
