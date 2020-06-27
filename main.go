package main

import (
	"api-simple/book"
	"api-simple/database"
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)

	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/book", book.GetBooks)
	v1.Get("/book/:id", book.GetBook)
	v1.Post("/book", book.NewBook)
	v1.Patch("/book/:id", book.UpdateBook)
	v1.Delete("/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DB, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	database.DB.AutoMigrate(&book.Book{})
	fmt.Println("Database Migrated")
}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	app.Listen(3000)

	defer database.DB.Close()
}
