package models

import "time"

type Project struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	ProjectURL  string    `json:"project_url"`
	CreatedAt   time.Time `json:"created_at"`
}