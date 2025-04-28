package controllers

import (
	"AtmGolang/database"
	"fmt"
	"log"
)

// Deposit menambah saldo akun
func Deposit(name string, Amount float64) {
	db := database.DB

	// Tambahkan saldo ke akun
	_, err := db.Exec("UPDATE accounts SET balance = balance + ? WHERE name = ?", Amount, name)
	if err != nil {
		log.Printf("Error saat deposit: %v", err)
		return
	}

	// Simpan riwayat transaksi
	_, err = db.Exec(`
		INSERT INTO transactions (IdAkun, Tipe, Amount, Target_id)
		VALUES ((SELECT id FROM accounts WHERE name = ?), 'Deposit', ?, NULL)
	`, name, Amount)
	if err != nil {
		log.Printf("Error saat mencatat transaksi deposit: %v", err)
		return
	}

	fmt.Printf("Deposit berhasil sebesar %.2f ke akun %s\n", Amount, name)
}

// Withdraw mengurangi saldo akun (tidak boleh negatif)
func Withdraw(name string, Amount float64) {
	db := database.DB

	// Cek saldo saat ini
	var balance float64
	err := db.QueryRow("SELECT balance FROM accounts WHERE name = ?", name).Scan(&balance)
	if err != nil {
		log.Printf("Error saat mengambil saldo: %v", err)
		return
	}

	if balance < Amount {
		fmt.Println("Saldo tidak cukup untuk penarikan.")
		return
	}

	// Update saldo
	_, err = db.Exec("UPDATE accounts SET balance = balance - ? WHERE name = ?", Amount, name)
	if err != nil {
		log.Printf("Error saat mengurangi saldo: %v", err)
		return
	}

	// Simpan riwayat transaksi
	_, err = db.Exec(`
		INSERT INTO transactions (IdAkun, Tipe, Amount, Target_id)
		VALUES ((SELECT id FROM accounts WHERE name = ?), 'Withdraw', ?, NULL)
	`, name, Amount)
	if err != nil {
		log.Printf("Error saat mencatat transaksi withdraw: %v", err)
		return
	}

	fmt.Printf("Penarikan sebesar %.2f berhasil!\n", Amount)
}

// Transfer mentransfer saldo ke akun lain
func Transfer(sender string, recipient string, Amount float64) {
	db := database.DB

	// Cek saldo pengirim
	var senderBalance float64
	err := db.QueryRow("SELECT balance FROM accounts WHERE name = ?", sender).Scan(&senderBalance)
	if err != nil {
		log.Printf("Error saat mengambil saldo pengirim: %v", err)
		return
	}

	if senderBalance < Amount {
		fmt.Println("Saldo tidak cukup untuk transfer.")
		return
	}

	// Kurangi saldo pengirim
	_, err = db.Exec("UPDATE accounts SET balance = balance - ? WHERE name = ?", Amount, sender)
	if err != nil {
		log.Printf("Error saat mengurangi saldo pengirim: %v", err)
		return
	}

	// Tambah saldo penerima
	_, err = db.Exec("UPDATE accounts SET balance = balance + ? WHERE name = ?", Amount, recipient)
	if err != nil {
		log.Printf("Error saat menambah saldo penerima: %v", err)
		return
	}

	// Simpan transaksi pengirim (Transfer Out)
	_, err = db.Exec(`
		INSERT INTO transactions (IdAkun, Tipe, Amount, Target_id)
		VALUES ((SELECT id FROM accounts WHERE name = ?), 'Transfer Out', ?, (SELECT id FROM accounts WHERE name = ?))
	`, sender, Amount, recipient)
	if err != nil {
		log.Printf("Error saat mencatat transaksi transfer out: %v", err)
		return
	}

	// Simpan transaksi penerima (Transfer In)
	_, err = db.Exec(`
		INSERT INTO transactions (IdAkun, Tipe, Amount, Target_id)
		VALUES ((SELECT id FROM accounts WHERE name = ?), 'Transfer In', ?, (SELECT id FROM accounts WHERE name = ?))
	`, recipient, Amount, sender)
	if err != nil {
		log.Printf("Error saat mencatat transaksi transfer in: %v", err)
		return
	}

	fmt.Printf("Transfer sebesar %.2f ke %s berhasil!\n", Amount, recipient)
}

// CekSaldo untuk melihat saldo akun

