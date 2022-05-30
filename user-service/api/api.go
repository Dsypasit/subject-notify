package api

import (
	"fmt"
	"log"
	"user-service/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetAccount(c *fiber.Ctx) error
	CreateAccount(c *fiber.Ctx) error
	UpdateInformation(c *fiber.Ctx) error
	UpdatePassword(c *fiber.Ctx) error
	DeleteAccount(c *fiber.Ctx) error
}

type userHandler struct {
	userSrv services.UserService
}

func NewUserHandler(userSrv services.UserService) UserHandler {
	return userHandler{userSrv}
}

func (h userHandler) GetAccount(c *fiber.Ctx) error {
	username := Username{}
	if err := c.BodyParser(&username); err != nil {
		return fiber.ErrExpectationFailed
	}
	user, err := h.userSrv.GetUser(username.Username)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h userHandler) CreateAccount(c *fiber.Ctx) error {
	user := services.NewUser{}
	if err := c.BodyParser(&user); err != nil {
		log.Println(err)
		return fiber.ErrExpectationFailed
	}
	userResponse, err := h.userSrv.CreateAccount(user)
	if err != nil {
		return err
	}
	return c.JSON(userResponse)
}

func (h userHandler) UpdateInformation(c *fiber.Ctx) error {
	user := services.UpdateUserInfo{}
	if err := c.BodyParser(&user); err != nil {
		return fiber.ErrExpectationFailed
	}
	if err := h.userSrv.UpdateInformation(user); err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "update informatin success",
	})
}

func (h userHandler) UpdatePassword(c *fiber.Ctx) error {
	user := services.UpdateUserPassword{}
	if err := c.BodyParser(&user); err != nil {
		return fiber.ErrExpectationFailed
	}
	if err := h.userSrv.UpdatePassword(user); err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": "update informatin success",
	})
}

func (h userHandler) DeleteAccount(c *fiber.Ctx) error {
	username := Username{}
	if err := c.BodyParser(&username); err != nil {
		return fiber.ErrExpectationFailed
	}
	if err := h.userSrv.DeleteAccount(username.Username); err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Account %v has been deleted", username.Username),
	})
}
