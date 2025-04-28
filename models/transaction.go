package models

// Enum untuk tipe transaksi
type TransactionType string

const (
    Deposit     TransactionType = "Deposit"
    Withdraw    TransactionType = "Withdraw"
    TransferIn  TransactionType = "TransferIn"
    TransferOut TransactionType = "TransferOut"
)

// Struct untuk menyimpan data transaksi
type Transaction struct {
    ID        int             `json:"id"`
    IdAkun    int             `json:"id_akun"`
    Tipe      TransactionType `json:"tipe"`
    Amount    float64         `json:"amount"`
    TargetID  int             `json:"target_id"`
}