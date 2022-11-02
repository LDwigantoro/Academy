package books

import (
	"github.com/LDwigantoro/go-fiber-api/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

func (h handler) DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, result.Error.Error())
	}

	// menghapus entri db
	h.DB.Delete(&book)

	return c.SendStatus(fiber.StatusOK)
}
