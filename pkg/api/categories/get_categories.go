package categories

import (
	"books/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetCategories(c *fiber.Ctx) error {
	var categories []models.Categories

	if result := h.DB.Preload("Book").Find(&categories); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&categories)
}
