package api

import (
	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App, handler UserHandler) {
	app.Post("/OpenAccount", handler.CreateAccount)
	app.Post("/GetAccount", handler.GetAccount)
	app.Post("/UpdateAccountInfo", handler.UpdateInformation)
	app.Post("/UpdateAccountPassword", handler.UpdatePassword)
	app.Post("/DeleteAccount", handler.DeleteAccount)
	app.Post("/Login", handler.Login)
	app.Get("/Hello", handler.Hello)
}
