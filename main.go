package main

import (
	userController "AtmGolang/controllers/user"
	transactionController "AtmGolang/controllers/transaction"
	"AtmGolang/database"
	"fmt"
	"os"
)

var isLoggedIn bool
var currentUser string

func main() {
	database.Connect()

	var choice int

	for {
		fmt.Println("\n==============================")
		fmt.Println("SELAMAT DATANG DI ATM GOLANG")
		fmt.Println("==============================")

		if !isLoggedIn {
			// Menu sebelum login
			fmt.Println("1. Registrasi User")
			fmt.Println("2. Login")
			fmt.Println("3. Exit")
			fmt.Print("Pilih menu: ")

			_, err := fmt.Scanln(&choice)
			if err != nil {
				fmt.Println("Gagal membaca input. Silakan coba lagi.")
				continue
			}

			switch choice {
			case 1:
				var name, pin string
				var balance float64

				fmt.Print("Masukkan nama: ")
				fmt.Scanln(&name)
				fmt.Print("Masukkan PIN: ")
				fmt.Scanln(&pin)
				fmt.Print("Masukkan saldo awal: ")
				fmt.Scanln(&balance)

				userController.CreateUserController(name, pin, balance)

			case 2:
				var name, pin string

				fmt.Print("Masukkan nama: ")
				fmt.Scanln(&name)
				fmt.Print("Masukkan PIN: ")
				fmt.Scanln(&pin)

				// Panggil Login
				success := userController.Login(name, pin)
				if success {
					isLoggedIn = true
					currentUser = name
					fmt.Println("Login berhasil!")
				} else {
					fmt.Println("Login gagal, nama atau PIN salah.")
				}

			case 3:
				fmt.Println("Terima kasih. Program selesai.")
				os.Exit(0)

			default:
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			}
		} else {
			// Menu setelah login
			fmt.Printf("\nSelamat datang, %s!\n", currentUser)
			fmt.Println("1. Cek Saldo")
			fmt.Println("2. Deposit")
			fmt.Println("3. Withdraw")
			fmt.Println("4. Transfer")
			fmt.Println("5. Logout")
			fmt.Print("Pilih menu: ")

			_, err := fmt.Scanln(&choice)
			if err != nil {
				fmt.Println("Gagal membaca input. Silakan coba lagi.")
				continue
			}

			switch choice {
			case 1:
				userController.CekSaldo(currentUser)

			case 2:
				var amount float64
				fmt.Print("Masukkan jumlah deposit: ")
				fmt.Scanln(&amount)
				transactionController.Deposit(currentUser, amount)

			case 3:
				var amount float64
				fmt.Print("Masukkan jumlah withdraw: ")
				fmt.Scanln(&amount)
				transactionController.Withdraw(currentUser, amount)

			case 4:
				var recipient string
				var amount float64
				fmt.Print("Masukkan nama penerima: ")
				fmt.Scanln(&recipient)
				fmt.Print("Masukkan jumlah transfer: ")
				fmt.Scanln(&amount)
				transactionController.Transfer(currentUser, recipient, amount)

			case 5:
				userController.Logout()
				isLoggedIn = false
				currentUser = ""
				fmt.Println("Logout berhasil.")

			default:
				fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			}
		}
	}
}
