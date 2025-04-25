package controllers

import (
	"AtmGolang/database"
	"fmt"
	"log"
)

// Fungsi untuk mengenkripsi PIN


// CreateUserController - Menangani pembuatan user baru
func CreateUserController(name string, pin string, balance float64) {
	db := database.DB
	// Query untuk memasukkan data user baru
	query := "INSERT INTO accounts (name, pin, balance) VALUES (?, ?, ?)"
	result, err := db.Exec(query, name, pin, balance)
	if err != nil {
		log.Printf("Error inserting user: %v", err)
		return
	}

	// Mengambil ID terakhir yang dimasukkan
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error fetching last insert ID: %v", err)
		return
	}

	fmt.Printf("User Berhasil Ditambahkan Dengan ID %d Dan Total Saldo %.2f\n", id, balance)
}
