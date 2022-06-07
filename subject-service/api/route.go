package api

import "github.com/gofiber/fiber/v2"

func Route(app *fiber.App, handler subjectHandler) {
	api := app.Group("/subject")
	api.Post("/create", handler.CreateSubject)
	api.Post("/delete", handler.DeleteSubject)
	api.Post("/update", handler.UpdateSubject)
	api.Get("/get", handler.GetSubjectsByQuery)
	api.Get("/get/:username", handler.GetSubjectsByUsername)
	app.Get("/Hello", handler.Hello)
}
