package main

import (
	"backend/database"
	"backend/handlers"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Muat .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Hubungkan ke Database
	database.ConnectDB()
	defer database.Pool.Close()

	app := fiber.New()

	// Konfigurasi CORS agar frontend bisa mengakses
	allowedOrigins := os.Getenv("CORS_ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		// Default jika tidak diset, berguna untuk development
		allowedOrigins = "http://localhost:5173"
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, HEAD, OPTIONS, PUT, DELETE", // Tambahkan metode lain jika perlu
		AllowCredentials: true,
	}))

	// Grup Rute API
	api := app.Group("/api")
	api.Get("/projects", handlers.GetProjects)

	log.Fatal(app.Listen(":3000"))
}