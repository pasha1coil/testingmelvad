package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/pasha1coil/testingmelvad/internal/models"
	"github.com/pasha1coil/testingmelvad/internal/service"
	log "github.com/sirupsen/logrus"
)

type Handlers struct {
	service   *service.TasksService
	validator *validator.Validate
}

func NewHandlers(service *service.TasksService) *Handlers {
	return &Handlers{
		service:   service,
		validator: validator.New(),
	}
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

	if err := h.validator.Struct(incr); err != nil {
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
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"value": value,
	})
}

// POST `/sign/hmacsha512`
func (h *Handlers) PostHmac(c *fiber.Ctx) error {
	log.Infoln("Start Handler - PostHmac")

	hmac := new(models.Hash)

	if err := c.BodyParser(hmac); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.validator.Struct(hmac); err != nil {
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
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"HMAC": value,
	})
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

	if err := h.validator.Struct(user); err != nil {
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
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"Id": id,
	})
}
