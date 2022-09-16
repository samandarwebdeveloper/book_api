package categories

import (
	"books/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	var category models.Categories

	if result := h.DB.Preload("Book").First(&category, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&category)
}
