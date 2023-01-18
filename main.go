package main

import (
	"github.com/gofiber/fiber/v2"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// make array of struct book
var books = []book{
	{ID: "1", Title: "In serach of Lost time", Author: "Mike Wills", Quantity: 2},
	{ID: "2", Title: "How times have changed", Author: "John Doe", Quantity: 10},
	{ID: "3", Title: "Interstellar", Author: "Dr.Hammerman", Quantity: 10},
	{ID: "4", Title: "African Woman", Author: "Lucy Wairimu", Quantity: 14},
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!")
	})

	app.Get("/books", func(c *fiber.Ctx) error {
		return c.JSON(books)
	})

	app.Get("/books/:id", func(c *fiber.Ctx) error {
		bookId := c.Params("id")

		for _, book := range books {
			if book.ID == bookId {
				return c.JSON(book)
			}
		}

		return c.SendString("No book found")
	})

	app.Post("/books", func(c *fiber.Ctx) error {
		newBook := new(book)

		if err := c.BodyParser(&newBook); err != nil {
			return err
		}

		//add book to books
		books = append(books, *newBook)

		return c.Status(fiber.StatusOK).JSON(books)

	})

	app.Listen(":8000")
}
