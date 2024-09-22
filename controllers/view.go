
// controllers/user.go
package controllers

import (
	"jwtFiber/models"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

var jwtSecret = []byte("your_secret")

// jwtGenerate generates a JWT token
func jwtGenerate(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Login authenticates the user and generates a JWT token
func Login(c *fiber.Ctx, db *gorm.DB) error {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	var user models.Users12
	if err := db.Where("username = ?", body.Username).First(&user).Error; err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
	}

	if user.Password != body.Password {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "The password is wrong"})
	}

	token, err := jwtGenerate(user.Username)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": token})
}

// Middleware to protect routes
func MiddleWare() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authorized := c.Get("Authorization")
		if authorized == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Authorization header not provided"})
		}

		tokenString := strings.TrimPrefix(authorized, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		return c.Next()
	}
}

// CreateUser creates a new user
func Create(c *fiber.Ctx, db *gorm.DB) error {
	var body models.Users12
	if err := c.BodyParser(&body); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

var existingUser models.Users12
if err := db.Where("username = ?", body.Username).First(&existingUser).Error; err == nil {
    return c.Status(http.StatusConflict).JSON(fiber.Map{"error": "Username already exists"})
}
	// Here, consider hashing the password before saving
	if res := db.Create(&body); res.Error != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": res.Error.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{"created": "User created successfully"})
}

// Protected route example
func Protected(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"message": "This is a protected route"})
}
