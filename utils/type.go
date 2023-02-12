package utils

type Result struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Message string `json:"message"`
	IsAdmin bool   `json:"is_admin"`
}
