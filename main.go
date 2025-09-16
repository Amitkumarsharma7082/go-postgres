package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// ! 4.
type Repository struct {
	// Repo is DB
	// gorm gives me to interact Database
	DB *gorm.DB
}

// ! 5.
type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

// ! 7.
// Create Books
func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}

	/*
		if you are using http : fiber is level of abstraction the background is
		request and response.
	*/
	// BodyParser : is Json you getting into book format
	err := context.BodyParser(&book)
	if err != nil {
		context.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Request failed"})
		return err
	}
	// added to database
	err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"})
		return err
	}
	context.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "book has been added"})

	return nil
}

// ! 8.
// GetBooks
func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.book{}

	// find books
	err := r.DB.Find(bookModels).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not get books"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "books fetched successfully",
		"data":    bookModels,
	})

	return nil
}

// ! 6.
// easily access the Repo inside func
// struct method due to *Repo
// the func accepts app *fiber
func (r *Repository) SetupRoutes(app *fiber.App) {
	// all the api is start from Api
	api := app.Group("/api")
	// CreateBook is method r *Repo
	api.Post("/create_books", r.CreateBook)
	api.Delete("delete/_book/:id", r.DeleteBook)
	api.Get("/get_books/:id", r.GetBookByID)
	api.Get("/books", r.GetBooks)
}

func main() {
	//! 1.
	// import .env file
	// capture the err
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// config from .env
	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("could not access the database")
	}
	r := Repository{
		DB: db,
	}
	//! 2.
	// use fiber package to create new routes
	app := fiber.New()
	//! 3.
	// create a func routes and app sent to it
	// here r : is Repository strcut
	r.SetupRoutes(app)

	// app listen on Port 8080
	app.Listen(":8080")
}
