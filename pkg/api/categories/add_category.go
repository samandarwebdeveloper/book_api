package categories

import (
	"books/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type AddCategoryRequestBody struct {
	Name      string `json:"name"`
}

func (h handler) AddCategory(c *fiber.Ctx) error {
	body := AddCategoryRequestBody{}

	// parse body, attach to AddCategoryRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var category models.Category

	category.Name = body.Name

	// insert new db entry
	if result := h.DB.Create(&category); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&category)
}
