package api

import (
	"log"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, handler UserHandler) {

	app.Post("/Login", handler.Login)
	app.Post("/OpenAccount", handler.CreateAccount)

	app.Use(Authorization)

	app.Get("/GetAccount", handler.GetAccount)
	app.Post("/UpdateAccountInfo", handler.UpdateInformation)
	app.Post("/UpdateAccountPassword", handler.UpdatePassword)
	app.Get("/DeleteAccount", handler.DeleteAccount)
	app.Get("/Logout", handler.Logout)
	app.Get("/Hello", handler.Hello)
}

func Authorization(c *fiber.Ctx) error {
	j := c.Cookies("j")
	a := c.Cookies("a")
	log.Printf("%#v %#v", j, a)
	if (j == "empty") || (a == "empty") {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}
	jwt_token := j + "." + a

	token, err := jwt.ParseWithClaims(jwt_token, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("kmutnb"), nil
	})
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "cant' parse with claims")
	}

	claims := token.Claims.(*jwt.StandardClaims)
	if token.Valid {
		c.Set("username", claims.Issuer)
	} else {
		return fiber.NewError(fiber.StatusUnauthorized, "token invalidate")
	}

	return c.Next()
}
