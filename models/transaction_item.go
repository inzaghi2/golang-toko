package models

type TransactionItem struct {
	ID            uint    `gorm:"primary_key" json:"id"`
	TransactionID uint    `json:"transaction_id"`
	ProductID     uint    `json:"product_id"`
	Quantity      int     `json:"quantity"`
	Price         float64 `json:"price"`
	// Tambahkan kolom-kolom lain sesuai kebutuhan
}
