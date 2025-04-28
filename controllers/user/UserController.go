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

func Login(name string, pin string) bool {
	db := database.DB
	// Query untuk memeriksa apakah user ada di database
	query := "SELECT id, name, pin, balance FROM accounts WHERE name = ? AND pin = ?"
	row := db.QueryRow(query, name, pin)

	var id int64
	var accountName string
	var accountPin string
	var balance float64

	err := row.Scan(&id, &accountName, &accountPin, &balance)
	if err != nil {
		log.Printf("Error fetching user: %v", err)
		return false
	}

	fmt.Printf("Login Berhasil! ID: %d, Name: %s, Balance: %.2f\n", id, accountName, balance)
	return true
}

func CekSaldo(name string) {
	db := database.DB

	var balance float64
	err := db.QueryRow("SELECT balance FROM accounts WHERE name = ?", name).Scan(&balance)
	if err != nil {
		log.Printf("Error saat mengambil saldo: %v", err)
		return
	}

	fmt.Printf("Saldo anda saat ini: %.2f\n", balance)
}


func Logout() {
	fmt.Println("Logout Berhasil!")
}

