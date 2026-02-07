// Package models menyimpan tipe data domain untuk aplikasi.
package models

import "time"

// Transaction merepresentasikan data transaksi pada sistem kasir.
type Transaction struct {
	ID          int                 `json:"id"`          // ID unik untuk transaksi.
	TotalAmount int                 `json:"total_amount"` // Total harga dari transaksi.
	CreatedAt   time.Time           `json:"created_at"`   // Waktu transaksi dibuat.
	Details     []TransactionDetail `json:"details"`      // Detail item dalam transaksi.
}

// TransactionDetail merepresentasikan detail item dalam satu transaksi.
type TransactionDetail struct {
	ID            int    `json:"id"`             // ID unik untuk detail transaksi.
	TransactionID int    `json:"transaction_id"` // ID transaksi yang terkait.
	ProductID     int    `json:"product_id"`     // ID produk yang dibeli.
	ProductName   string `json:"product_name,omitempty"` // Nama produk (opsional, dari join).
	Quantity      int    `json:"quantity"`       // Jumlah barang yang dibeli.
	Subtotal      int    `json:"subtotal"`       // Subtotal harga (harga * quantity).
}

// CheckoutItem merepresentasikan item yang akan di-checkout.
type CheckoutItem struct {
	ProductID int `json:"product_id"` // ID produk yang dibeli.
	Quantity  int `json:"quantity"`   // Jumlah barang yang dibeli.
}

// CheckoutRequest merepresentasikan request body untuk checkout.
type CheckoutRequest struct {
	Items []CheckoutItem `json:"items"` // Daftar item yang akan dibeli.
}
