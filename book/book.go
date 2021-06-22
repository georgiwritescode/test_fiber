package book

import (
	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	return c.SendString("Get all books")
}

func GetBook(c *fiber.Ctx) error {
	return c.SendString("Get a book!")
}

func CreateBook(c *fiber.Ctx) error {
	return c.SendString("Update a book!")
}

func DeleteBook(c *fiber.Ctx) error {
	return c.SendString("Delete a book!")
}
