package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	domain "github.com/qchart-app/service-tv-udf/internal/domain"
	"github.com/qchart-app/service-tv-udf/pkg/util"
)

type UserHandler struct {
	userUsecase domain.UserUseCase
}

func NewUserHandler(userUsecase domain.UserUseCase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api/v1/users")

	api.Post("/", h.createUser)
	api.Get("/:id", h.getUserByID)
	api.Put("/:id", h.updateUser)
	api.Delete("/:id", h.deleteUser)
}

func (h *UserHandler) createUser(c *fiber.Ctx) error {
	var user domain.User
	err := c.BodyParser(&user)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request body")
	}

	err = h.userUsecase.CreateUser(c.Context(), &user)
	if err != nil {
		return errors.Wrap(err, "failed to create user")
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "user created successfully",
		"user":    user,
	})
}

func (h *UserHandler) getUserByID(c *fiber.Ctx) error {
	id, err := util.ToUint(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.userUsecase.GetUserByID(c.Context(), id)
	if err != nil {
		return fiber.ErrNotFound
	}

	return c.JSON(user)
}

func (h *UserHandler) updateUser(c *fiber.Ctx) error {
	id, err := util.ToInt(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return fiber.ErrBadRequest
	}

	user.ID = id

	if err := h.userUsecase.UpdateUser(c.Context(), &user); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(user)
}

func (h *UserHandler) deleteUser(c *fiber.Ctx) error {
	id, err := util.ToUint(c.Params("id"))
	if err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.userUsecase.DeleteUser(c.Context(), id); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(fiber.StatusNoContent)
}
