package api

import (
	"fmt"
	"log"
	"strings"
	"time"
	"user-service/services"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetAccount(c *fiber.Ctx) error
	CreateAccount(c *fiber.Ctx) error
	UpdateInformation(c *fiber.Ctx) error
	UpdatePassword(c *fiber.Ctx) error
	DeleteAccount(c *fiber.Ctx) error
	Hello(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type userHandler struct {
	userSrv services.UserService
}

const secreate = "kmutnb"

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
	log.Printf("%#v", user)
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

func (h userHandler) Hello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello world",
		"status":  "ok",
	})
}

func (h userHandler) Login(c *fiber.Ctx) error {
	login := services.LoginUser{}
	if err := c.BodyParser(&login); err != nil {
		return fiber.ErrExpectationFailed
	}
	if err := h.userSrv.Login(login); err != nil {
		return err
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    "user",
		ExpiresAt: jwt.At(time.Now().Add(time.Hour * 24)),
	})

	token, err := claims.SignedString([]byte(secreate))
	if err != nil {
		return fiber.ErrExpectationFailed
	}

	split_token := strings.Split(token, ".")

	cookie_j := new(fiber.Cookie)
	cookie_j.Name = "j"
	cookie_j.Value = split_token[0] + "." + split_token[1]
	cookie_j.Secure = true
	cookie_j.HTTPOnly = true
	c.Cookie(cookie_j)

	cookie_a := new(fiber.Cookie)
	cookie_a.Name = "a"
	cookie_a.Value = split_token[2]
	cookie_a.Secure = true
	cookie_a.HTTPOnly = true
	c.Cookie(cookie_a)

	return c.JSON(fiber.Map{
		"message": "login success",
		"status":  "ok",
	})
}
