package handlers

import (
	"context"
	"log"
	"backend/database"
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetProjects(c *fiber.Ctx) error {
	rows, err := database.Pool.Query(context.Background(), "SELECT id, title, description, image_url, project_url, created_at FROM projects ORDER BY created_at DESC")
	if err != nil {
		log.Printf("Query failed: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "database error"})
	}
	defer rows.Close()

	projects := make([]models.Project, 0)
	for rows.Next() {
		var p models.Project
		if err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.ImageURL, &p.ProjectURL, &p.CreatedAt); err != nil {
			log.Printf("Scan failed: %v\n", err)
			continue
		}
		projects = append(projects, p)
	}

	return c.Status(fiber.StatusOK).JSON(projects)
}