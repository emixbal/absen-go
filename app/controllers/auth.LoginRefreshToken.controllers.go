package controllers

import (
	"absen-go/app/models"
	"absen-go/app/requests"
	"absen-go/config"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/golang-jwt/jwt"
	"github.com/gookit/validate"
	"golang.org/x/crypto/bcrypt"
)

var refreshSecret = []byte(os.Getenv("REFRESH_SECRET"))

func LoginWithRefrehToken(c *fiber.Ctx) error {
	var userClaim models.UserClaim

	p := new(requests.LoginForm)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	v := validate.Struct(p)
	if !v.Validate() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": v.Errors.One(),
		})
	}
	user := new(models.User)

	db := config.GetDBInstance()

	if res := db.Preload("Role").Where("email = ?", p.Email).First(&user); res.RowsAffected <= 0 {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid Email!",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password)); err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error":   true,
			"message": "Password is incorrect!",
		})
	}
	userClaim.Issuer = utils.UUIDv4()
	userClaim.Id = int(user.ID)
	userClaim.Email = user.Email
	userClaim.Role = user.Role.Name
	accessToken, refreshToken := models.GenerateTokens(&userClaim, false)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func RefreshToken(c *fiber.Ctx) error {
	var userClaim models.UserClaim

	p := new(requests.RefreshTokenForm)
	if err := c.BodyParser(p); err != nil {
		log.Println(err)
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Empty payloads",
		})
	}
	v := validate.Struct(p)
	if !v.Validate() {
		log.Println(v.Errors.One())
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message":      v.Errors.One(),
			"RefreshToken": p.RefreshToken,
		})
	}

	refreshToken := p.RefreshToken
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return refreshSecret, nil
	})

	if err != nil {
		log.Println("the error from parse: ", err)
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": err,
		})
	}

	//is token valid?
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "StatusUnauthorized",
		})
	}

	user_id := claims["user_id"]
	user_id_int := int(user_id.(float64))
	userClaim.Issuer = fmt.Sprintf("%v", claims["issuer"])
	userClaim.Id = user_id_int
	userClaim.Email = fmt.Sprintf("%v", claims["email"])
	userClaim.Role = fmt.Sprintf("%v", claims["role"])

	// if fail refresh token
	accessToken, refreshToken := models.GenerateTokens(&userClaim, true)
	if len(accessToken) < 1 || len(refreshToken) < 1 {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "something went wrong",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
