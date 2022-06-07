package errs

import "github.com/gofiber/fiber/v2"

var (
	InvalidRequest    = fiber.NewError(fiber.StatusBadRequest, "Invalid request")
	DuplicateUsername = fiber.NewError(fiber.StatusBadRequest, "Duplicate username")
	ServerError       = fiber.NewError(fiber.StatusInternalServerError, "database error")
	NotFound          = fiber.NewError(fiber.StatusNotFound, "Not Found")
	UserNotFound      = fiber.NewError(fiber.StatusNotFound, "Incorrect username")
	PasswordInvalid   = fiber.NewError(fiber.StatusNotFound, "Invalid password")
)
