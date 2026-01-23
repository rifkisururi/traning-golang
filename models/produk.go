// Package models menyimpan tipe data domain untuk aplikasi.
package models

// Produk merepresentasikan data produk pada sistem kasir.
type Produk struct {
	ID    int    `json:"id"`    // ID unik untuk produk.
	Nama  string `json:"nama"`  // Nama produk yang tampil di API.
	Harga int    `json:"harga"` // Harga produk dalam satuan rupiah.
	Stok  int    `json:"stok"`  // Stok tersedia untuk produk ini.
}
