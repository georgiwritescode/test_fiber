package main

import (
	"github.com/georgiwritescode/test_fiber/book"
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var (
	db  *gorm.DB
	err error
)

func setup(app *fiber.App) {
	app.Get("/", greeting)
	app.Get("/api/v1/books", book.GetBooks)
	app.Get("/api/v1/books/:id", book.GetBook)
	app.Post("/api/v1/books", book.CreateBook)
	app.Delete("/api/v1/books/:id", book.DeleteBook)

}

//TODO connect to db
func connect2db() {
	db, err = gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres sslmode=disable password=postgres")

	if err != nil {
		log.Fatal("failed to connect database")
	}

	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal("[ERROR] Connection to db failed!")
		}
	}(db)

}

func main() {

	app := fiber.New()

	setup(app)

	err := app.Listen(":3000")
	if err != nil {
		log.Print("[ERROR] Failed to listen to port 3000")
	}

	connect2db()
}

func greeting(c *fiber.Ctx) error {
	return c.SendString("Hello there!")
}
