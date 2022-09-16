package books

import (
	"books/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func (h handler) GetBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book

	if result := h.DB.Preload("Category").First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusOK).JSON(&book)
}
