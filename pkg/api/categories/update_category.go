package categories

import (
	"books/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type UpdateCategoryRequestBody struct {
	Name string `json:"name"`
}

func (h handler) UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateCategoryRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var category models.Category

	if result := h.DB.First(&category, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	category.Name = body.Name

	// save book
	h.DB.Save(&category)

	return c.Status(fiber.StatusOK).JSON(&category)
}
