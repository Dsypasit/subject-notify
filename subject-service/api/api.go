package api

import (
	"fmt"
	"subject-service/errs"
	"subject-service/repositories"
	"subject-service/services"

	"github.com/gofiber/fiber/v2"
)

type subjectHandler struct {
	subjectSrv services.SubjectService
}

type updateSubjectReq struct {
	Query  repositories.SubjectQuery
	Update repositories.SubjectUpdate
}

func NewSubjectHandler(subjectSrv services.SubjectService) subjectHandler {
	return subjectHandler{subjectSrv}
}

func (h subjectHandler) CreateSubject(c *fiber.Ctx) error {
	reqSub := repositories.Subject{}
	if err := c.BodyParser(&reqSub); err != nil {
		return errs.ServerError
	}
	if err := h.subjectSrv.CreateSubject(reqSub); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": fmt.Sprintf("Created %v subject by %v", reqSub.SubjectName, reqSub.Username),
	})
}

func (h subjectHandler) DeleteSubject(c *fiber.Ctx) error {
	reqSub := repositories.SubjectQuery{}
	if err := c.BodyParser(&reqSub); err != nil {
		return errs.ServerError
	}
	if err := h.subjectSrv.DeleteSubject(reqSub); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": fmt.Sprintf("Deleted %v subject by %v", reqSub.SubjectName, reqSub.Username),
	})
}

func (h subjectHandler) UpdateSubject(c *fiber.Ctx) error {
	reqSub := updateSubjectReq{}
	if err := c.BodyParser(&reqSub); err != nil {
		return errs.ServerError
	}

	subQuery := reqSub.Query
	subUpdate := reqSub.Update

	if err := h.subjectSrv.UpdateSubject(subQuery, subUpdate); err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": fmt.Sprintf("Updated %v subject by %v", subQuery.SubjectName, subQuery.Username),
	})
}

func (h subjectHandler) GetSubjectsByUsername(c *fiber.Ctx) error {
	reqSub := repositories.SubjectQuery{}
	if err := c.BodyParser(&reqSub); err != nil {
		return errs.ServerError
	}
	subjects, err := h.subjectSrv.GetSubjectsByUsername(reqSub)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(subjects)
}

func (h subjectHandler) GetSubjectsByQuery(c *fiber.Ctx) error {
	reqSub := repositories.SubjectQuery{}
	if err := c.BodyParser(&reqSub); err != nil {
		return errs.ServerError
	}
	subjects, err := h.subjectSrv.GetSubjectsByQuery(reqSub)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(subjects)
}
