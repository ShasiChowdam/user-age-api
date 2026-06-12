package handler

import (
	"github.com/ShasiChowdam/user-age-api/internal/service"
	"github.com/ShasiChowdam/user-age-api/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"

	"strconv"
)

type UserHandler struct {
	service  *service.UserService
	validate *validator.Validate
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service:  service,
		validate: validator.New(),
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := h.service.CreateUser(
		c.Context(),
		req.Name,
		req.DOB,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB: user.Dob.Time.Format("2006-01-02"),
	})
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	user, age, err := h.service.GetUserByID(
		c.Context(),
		int32(id),
	)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(models.UserWithAgeResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Time.Format("2006-01-02"),
		Age:  age,
	})
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.service.ListUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(users)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	var req models.UpdateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := h.service.UpdateUser(
		c.Context(),
		int32(id),
		req.Name,
		req.DOB,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(models.UserResponse{
		ID:   user.ID,
		Name: user.Name,
		DOB:  user.Dob.Time.Format("2006-01-02"),
	})
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid user id",
		})
	}

	err = h.service.DeleteUser(
		c.Context(),
		int32(id),
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}