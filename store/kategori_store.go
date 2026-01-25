package store

import (
	"traning-golang/models"
	"sync"
)

// in-memory storage untuk kategori
var (
	kategori     []models.Kategori
	kategoriNextID int = 1
	kategoriMutex   sync.RWMutex
)

// init menginisialisasi data dummy kategori
func init() {
	kategori = []models.Kategori{
		{ID: 1, Nama: "Makanan", Deskripsi: "Kategori untuk produk makanan"},
		{ID: 2, Nama: "Minuman", Deskripsi: "Kategori untuk produk minuman"},
		{ID: 3, Nama: "Bumbu Dapur", Deskripsi: "Kategori untuk bumbu dan rempah"},
	}
	kategoriNextID = 4
}

// GetAllKategori mengembalikan semua kategori
func GetAllKategori() []models.Kategori {
	kategoriMutex.RLock()
	defer kategoriMutex.RUnlock()

	// Buat copy untuk mencegah mutation dari luar
	result := make([]models.Kategori, len(kategori))
	copy(result, kategori)
	return result
}

// GetKategoriByID mencari kategori berdasarkan ID
func GetKategoriByID(id int) (models.Kategori, bool) {
	kategoriMutex.RLock()
	defer kategoriMutex.RUnlock()

	for _, k := range kategori {
		if k.ID == id {
			return k, true
		}
	}
	return models.Kategori{}, false
}

// GetKategoriNextID mengembalikan ID berikutnya untuk kategori baru
func GetKategoriNextID() int {
	kategoriMutex.Lock()
	defer kategoriMutex.Unlock()

	id := kategoriNextID
	kategoriNextID++
	return id
}

// AddKategori menambahkan kategori baru
func AddKategori(k models.Kategori) {
	kategoriMutex.Lock()
	defer kategoriMutex.Unlock()

	k.ID = GetKategoriNextID()
	kategori = append(kategori, k)
}

// UpdateKategori mengupdate kategori yang sudah ada
func UpdateKategori(id int, updated models.Kategori) bool {
	kategoriMutex.Lock()
	defer kategoriMutex.Unlock()

	for i, k := range kategori {
		if k.ID == id {
			updated.ID = id
			kategori[i] = updated
			return true
		}
	}
	return false
}

// DeleteKategori menghapus kategori berdasarkan ID
func DeleteKategori(id int) bool {
	kategoriMutex.Lock()
	defer kategoriMutex.Unlock()

	for i, k := range kategori {
		if k.ID == id {
			// Delete dengan menghapus elemen dari slice
			kategori = append(kategori[:i], kategori[i+1:]...)
			return true
		}
	}
	return false
}
