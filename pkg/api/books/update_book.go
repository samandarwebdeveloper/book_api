package books

import (
	"books/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type UpdateBookRequestBody struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	Stock      int32  `json:"stock"`
	Price      int32  `json:"price"`
	CategoryId int32  `json:"category_id"`
}

func (h handler) UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")
	body := UpdateBookRequestBody{}

	// getting request's body
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	book.Title = body.Title
	book.Author = body.Author
	book.Stock = body.Stock
	book.Price = body.Price
	book.CategoryId = body.CategoryId

	// save book
	h.DB.Save(&book)

	return c.Status(fiber.StatusOK).JSON(&book)
}
