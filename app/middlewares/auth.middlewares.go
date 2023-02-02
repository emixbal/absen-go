package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var (
	jwtKey = []byte(os.Getenv("JWT_SECRET"))
)

func IsAuthenticated(c *fiber.Ctx) error {
	raw_token := c.Request().Header.Peek("Authorization")
	tokenString := string(raw_token)

	if tokenString == "" {
		return c.Status(http.StatusForbidden).JSON(
			map[string]string{
				"message": "Unauthorized, need access token to access this API route!",
			},
		)
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusUnauthorized).JSON(
			map[string]string{
				"message": "token expired!",
			},
		)
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		c.Locals("user_id", claims["user_id"])
		c.Locals("user_email", claims["email"])
		c.Locals("role", claims["role"])

		return c.Next()
	}

	return c.Status(http.StatusForbidden).JSON(
		map[string]string{
			"message": "Unauthorized, access token is invalid!",
		},
	)
}

func IsAdmin(c *fiber.Ctx) error {
	if c.Locals("role") == "admin" {
		return c.Next()
	}

	if c.Locals("role") == "superadmin" {
		return c.Next()
	}

	return c.Status(http.StatusForbidden).JSON(
		map[string]string{
			"message": "Unauthorized to access this menu",
		},
	)
}

func IsSuperadmin(c *fiber.Ctx) error {
	if c.Locals("role") == "superadmin" {
		return c.Next()
	}

	return c.Status(http.StatusForbidden).JSON(
		map[string]string{
			"message": "Unauthorized to access this menu",
		},
	)
}
