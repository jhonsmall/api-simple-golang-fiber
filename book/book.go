package book

import (
	"api-simple/database"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	db := database.DB
	var books []Book
	db.Find(&books)
	c.JSON(books)
}

func GetBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB
	var book Book
	db.Find(&book, id)
	c.JSON(book)
}

func NewBook(c *fiber.Ctx) {
	db := database.DB
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&book)
	c.JSON(book)
}

func UpdateBook(c *fiber.Ctx) {
	type DataUpdateBook struct {
		Title  string `json:"title"`
		Author string `json:"author"`
		Rating int    `json:"rating"`
	}
	var dataUB DataUpdateBook
	if err := c.BodyParser(&dataUB); err != nil {
		c.Status(503).Send(err)
		return
	}
	var book Book
	id := c.Params("id")
	db := database.DB
	db.First(&book, id)

	book = Book{
		Title:  dataUB.Title,
		Author: dataUB.Author,
		Rating: dataUB.Rating,
	}
	db.Save(&book)
	c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DB

	var book Book
	db.First(&book, id)
	if book.Title == "" {
		c.Status(500).Send("No Book Found with ID")
		return
	}
	db.Delete(&book)
	c.Send("Book Successfully deleted")
}
