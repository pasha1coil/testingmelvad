package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
	"testingMelvad/internal/models"
	"testingMelvad/internal/service"
)

type Handlers struct {
	service *service.TasksService
}

func NewHandlers(service *service.TasksService) *Handlers {
	return &Handlers{service: service}
}

// POST `/redis/incr`
func (h *Handlers) PostIncr(c *fiber.Ctx) error {
	log.Infoln("Start Handler - PostIncr")

	incr := new(models.Incr)

	if err := c.BodyParser(incr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(incr); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	value, err := h.service.Increment(incr)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(value)
}

// POST `/sign/hmacsha512`
func (h *Handlers) PostHmac(c *fiber.Ctx) error {
	log.Infoln("Start Handler - PostHmac")

	hmac := new(models.Hmac)

	if err := c.BodyParser(hmac); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(hmac); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	value, err := h.service.CalculateHMAC(hmac)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(value)
}

// POST `/postgres/users`
func (h *Handlers) PostUsers(c *fiber.Ctx) error {
	log.Infoln("Start Handler - PostUsers")

	user := new(models.Users)

	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	id, err := h.service.AddUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(id)
}
