package main

import (
	"AtmGolang/controllers"
	"AtmGolang/database"
	"fmt"
	"os"
)

func main() {
	database.Connect() 
	var choice int
	for {
		// Menampilkan menu pilihan
		fmt.Println("\nSELAMAT DATANG DI ATM GOLANG")
		fmt.Println("1. Registrasi User")
		fmt.Println("2. Exit")
		fmt.Print("Pilih menu: ")

		// Membaca input dari terminal
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Gagal membaca input. Silakan coba lagi.")
			continue
		}

		switch choice {
		case 1:
			var name, pin string
			var balance float64

			// Memasukkan data untuk user baru
			fmt.Print(" Masukan nama: ")
			fmt.Scanln(&name)
			fmt.Print("Masukan pin: ")
			fmt.Scanln(&pin)
			fmt.Print("Tentukan Jumlah: ")
			fmt.Scanln(&balance)

			// Memanggil controller untuk membuat user
			controllers.CreateUserController(name, pin, balance)

		case 2:
			// Exit
			fmt.Println("Exiting program...")
			os.Exit(0)

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
