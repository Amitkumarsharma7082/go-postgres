package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Repository struct {
	// Repo is DB
	// gorm gives me to interact Database
	DB *gorm.DB
}
type Book struct {
	Author    string `json:"author"`
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
}

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

	// use fiber package to create new routes
	app := fiber.New()
	// create a func routes and app sent to it
	// here r : is Repository strcut
	r.SetupRoutes(app)
	// app listen on Port 8080
	app.Listen(":8080")
}
