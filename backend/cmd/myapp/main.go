package main

import (
	"backend/internal/database"
	"backend/internal/handlers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	database.ConnectDB()
	defer database.Pool.Close()

	app := fiber.New()

	allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:5173"
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, HEAD, OPTIONS, PUT, DELETE",
		AllowCredentials: true,
	}))

	// Grup Rute API
	api := app.Group("/api")
	api.Get("/projects", handlers.GetProjects)

	log.Fatal(app.Listen(":3000"))
}