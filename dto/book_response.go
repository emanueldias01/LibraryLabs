package dto

type BookResponse struct {
	ID uint `json"id"`
	Name string `json:"name"`
	Author string `json:"author"`
	YearPublication uint `json:"year_publication"`
	Publisher string `json:"publisher"`
	PagesNumber uint `json:"pages_number"`
	Language string `json:"language"`
}