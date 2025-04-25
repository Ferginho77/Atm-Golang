package models

// User struct untuk menyimpan data user
type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Pin      string  `json:"pin"`
	Balance  float64 `json:"balance"`
}
