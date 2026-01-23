// Package store menyimpan data sementara (in-memory) untuk aplikasi.
package store

import (
	"log"

	"kasir-api/models"
)

// produk adalah penyimpanan data produk di memori (sementara).
var produk = []models.Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
	{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
	{ID: 3, Nama: "kecap", Harga: 12000, Stok: 20},
}

// GetAll mengembalikan semua data produk.
func GetAll() []models.Produk {
	// Log alur data saat mengambil semua produk.
	log.Printf("[flow-store-1] GetAll total=%d", len(produk))

	// Buat salinan agar slice internal tidak bisa diubah langsung.
	data := make([]models.Produk, len(produk))
	copy(data, produk)
	return data
}

// GetByID mengembalikan satu produk berdasarkan ID.
func GetByID(id int) (models.Produk, bool) {
	// Log alur data saat mencari produk tertentu.
	log.Printf("[flow-store-1] GetByID start id=%d total=%d", id, len(produk))
	for _, p := range produk {
		if p.ID == id {
			log.Printf("[flow-store-2] GetByID found id=%d", p.ID)
			return p, true
		}
	}
	log.Printf("[flow-store-2] GetByID not found id=%d", id)
	return models.Produk{}, false
}

// NextID membuat ID baru berdasarkan panjang data saat ini.
func NextID() int {
	// Log perhitungan ID baru.
	next := len(produk) + 1
	log.Printf("[flow-store-1] NextID total=%d next_id=%d", len(produk), next)
	return next
}

// Add menambahkan produk baru ke penyimpanan.
func Add(p models.Produk) models.Produk {
	// Log alur data saat menambah produk.
	log.Printf("[flow-store-1] Add before_total=%d id=%d", len(produk), p.ID)
	produk = append(produk, p)
	log.Printf("[flow-store-2] Add after_total=%d id=%d", len(produk), p.ID)
	return p
}

// Update mengganti data produk berdasarkan ID.
func Update(id int, p models.Produk) (models.Produk, bool) {
	// Log alur data saat update produk.
	log.Printf("[flow-store-1] Update start id=%d", id)
	for i := range produk {
		if produk[i].ID == id {
			log.Printf("[flow-store-2] Update match id=%d", id)
			p.ID = id
			produk[i] = p
			log.Printf("[flow-store-3] Update saved id=%d", id)
			return p, true
		}
	}
	log.Printf("[flow-store-2] Update not found id=%d", id)
	return models.Produk{}, false
}

// Delete menghapus produk berdasarkan ID.
func Delete(id int) bool {
	// Log alur data saat menghapus produk.
	log.Printf("[flow-store-1] Delete start id=%d total=%d", id, len(produk))
	for i, p := range produk {
		if p.ID == id {
			log.Printf("[flow-store-2] Delete match id=%d index=%d", p.ID, i)
			produk = append(produk[:i], produk[i+1:]...)
			log.Printf("[flow-store-3] Delete done total=%d", len(produk))
			return true
		}
	}
	log.Printf("[flow-store-2] Delete not found id=%d", id)
	return false
}
