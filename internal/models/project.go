package models

type Project struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Image       string   `json:"image"`
	Techstack   []string `json:"techstack"`
}
