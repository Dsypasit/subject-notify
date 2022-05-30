package errs

import "github.com/gofiber/fiber/v2"

var (
	InvalidRequest    = fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	DuplicateUsername = fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	ServerError       = fiber.NewError(fiber.StatusInternalServerError, "database error")
	NotFound          = fiber.NewError(fiber.StatusNotFound, "Not Found")
)
