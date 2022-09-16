package books

import (
	"books/pkg/models"

	"github.com/gofiber/fiber/v2"
)

type AddBookRequestBody struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	Stock      int32  `json:"stock"`
	Price      int32  `json:"price"`
	CategoryId int32  `json:"category_id"`
}

func (h handler) AddBook(c *fiber.Ctx) error {
	body := AddBookRequestBody{}

	// parse body, attach to AddBookRequestBody struct
	if err := c.BodyParser(&body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var book models.Book

	book.Title = body.Title
	book.Author = body.Author
	book.Stock = body.Stock
	book.Price = body.Price
	book.CategoryId = body.CategoryId

	// insert new db entry
	if result := h.DB.Create(&book); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&book)
}
