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

	//display hello world
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!")
	})

	//get all books
	app.Get("/books", func(c *fiber.Ctx) error {
		return c.JSON(books)
	})

	//get book
	app.Get("/books/:id", func(c *fiber.Ctx) error {
		bookId := c.Params("id")

		// for _, book := range books {
		// 	if book.ID == bookId {
		// 		return c.JSON(book)
		// 	}
		// }
		book, err := getBookById(bookId)

		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(book)
	})

	//create book
	app.Post("/books", func(c *fiber.Ctx) error {
		newBook := new(book)

		if err := c.BodyParser(&newBook); err != nil {
			return err
		}

		//add book to books
		books = append(books, *newBook)

		return c.Status(fiber.StatusOK).JSON(books)

	})

	//checkout book
	app.Patch("/checkout", func(c *fiber.Ctx) error {
		//get query value from the url
		id := c.Query("id")

		//get book
		book, err := getBookById(id)

		//check error
		if err != nil {
			return err
		}

		//check quantity of book
		//if less than zero cannot checkout
		if book.Quantity <= 0 {
			return c.Status(fiber.StatusBadRequest).SendString("Cannot checkout book")
		}

		//decrease quantity by 1
		book.Quantity -= 1

		//send status and book
		return c.Status(fiber.StatusOK).JSON(book)
	})

	//return book func
	app.Patch("/return", func(c *fiber.Ctx) error {
		//get id from query
		id := c.Query("id")

		//get book
		book, err := getBookById(id)

		//check err
		if err != nil {
			return err
		}

		//increase book quantity by 1
		book.Quantity += 1

		//send status and book
		return c.Status(fiber.StatusOK).JSON(book)
	})

	app.Listen(":8000")
}

// get book by id
func getBookById(id string) (*book, error) {

	for i, book := range books {
		if book.ID == id {
			//pointer to the memory address and return error as nil
			return &books[i], nil
		}
	}

	//return nil books and error
	return nil, fiber.NewError(400, "Book Not Found")
}
