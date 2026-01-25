package models

// Kategori merepresentasikan model kategori produk
type Kategori struct {
	ID          int    `json:"id"`          // Unique ID untuk kategori
	Nama        string `json:"nama"`        // Nama kategori
	Deskripsi   string `json:"deskripsi"`   // Deskripsi kategori
}
