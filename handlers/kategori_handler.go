package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"traning-golang/models"
	"traning-golang/store"
)

// GetKategoriHandler mengambil semua kategori
func GetKategoriHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[flow-1] Masuk ke GetKategoriHandler")

	w.Header().Set("Content-Type", "application/json")

	log.Println("[flow-2] Mengambil data dari store")
	kategori := store.GetAllKategori()

	log.Println("[flow-3] Mengencode data ke JSON dan mengirim response")
	json.NewEncoder(w).Encode(kategori)
	log.Println("[flow-4] Selesai mengirim response")
}

// GetKategoriByIDHandler mengambil kategori berdasarkan ID
func GetKategoriByIDHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[flow-1] Masuk ke GetKategoriByIDHandler")

	w.Header().Set("Content-Type", "application/json")

	// Ambil ID dari URL parameter
	idStr := r.URL.Path[len("/api/kategori/"):]
	log.Printf("[flow-2] Mengambil ID dari URL: %s", idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[flow-3] Error: ID tidak valid - %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "ID tidak valid",
		})
		return
	}

	log.Printf("[flow-4] Mencari kategori dengan ID: %d", id)
	kategori, found := store.GetKategoriByID(id)

	if !found {
		log.Printf("[flow-5] Kategori dengan ID %d tidak ditemukan", id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf("Kategori dengan ID %d tidak ditemukan", id),
		})
		return
	}

	log.Printf("[flow-6] Kategori ditemukan: %s", kategori.Nama)
	json.NewEncoder(w).Encode(kategori)
	log.Println("[flow-7] Selesai mengirim response")
}

// CreateKategoriHandler menambahkan kategori baru
func CreateKategoriHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[flow-1] Masuk ke CreateKategoriHandler")

	// Pastikan method adalah POST
	if r.Method != http.MethodPost {
		log.Printf("[flow-2] Method tidak diizinkan: %s", r.Method)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Method tidak diizinkan",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")

	log.Println("[flow-3] Decode request body")
	var kategori models.Kategori
	err := json.NewDecoder(r.Body).Decode(&kategori)
	if err != nil {
		log.Printf("[flow-4] Error decoding JSON: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Format JSON tidak valid",
		})
		return
	}

	// Validasi input
	if kategori.Nama == "" {
		log.Println("[flow-5] Error: Nama kategori tidak boleh kosong")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Nama kategori tidak boleh kosong",
		})
		return
	}

	log.Printf("[flow-6] Menambahkan kategori baru: %s", kategori.Nama)
	store.AddKategori(kategori)

	log.Println("[flow-7] Kategori berhasil ditambahkan")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Kategori berhasil ditambahkan",
	})
}

// UpdateKategoriHandler mengupdate kategori yang sudah ada
func UpdateKategoriHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[flow-1] Masuk ke UpdateKategoriHandler")

	// Pastikan method adalah PUT
	if r.Method != http.MethodPut {
		log.Printf("[flow-2] Method tidak diizinkan: %s", r.Method)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Method tidak diizinkan",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Ambil ID dari URL parameter
	idStr := r.URL.Path[len("/api/kategori/"):]
	log.Printf("[flow-3] Mengambil ID dari URL: %s", idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[flow-4] Error: ID tidak valid - %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "ID tidak valid",
		})
		return
	}

	log.Println("[flow-5] Decode request body")
	var kategori models.Kategori
	err = json.NewDecoder(r.Body).Decode(&kategori)
	if err != nil {
		log.Printf("[flow-6] Error decoding JSON: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Format JSON tidak valid",
		})
		return
	}

	// Validasi input
	if kategori.Nama == "" {
		log.Println("[flow-7] Error: Nama kategori tidak boleh kosong")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Nama kategori tidak boleh kosong",
		})
		return
	}

	log.Printf("[flow-8] Mengupdate kategori dengan ID: %d", id)
	success := store.UpdateKategori(id, kategori)

	if !success {
		log.Printf("[flow-9] Kategori dengan ID %d tidak ditemukan", id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf("Kategori dengan ID %d tidak ditemukan", id),
		})
		return
	}

	log.Println("[flow-10] Kategori berhasil diupdate")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Kategori berhasil diupdate",
	})
}

// DeleteKategoriHandler menghapus kategori
func DeleteKategoriHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("[flow-1] Masuk ke DeleteKategoriHandler")

	// Pastikan method adalah DELETE
	if r.Method != http.MethodDelete {
		log.Printf("[flow-2] Method tidak diizinkan: %s", r.Method)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "Method tidak diizinkan",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Ambil ID dari URL parameter
	idStr := r.URL.Path[len("/api/kategori/"):]
	log.Printf("[flow-3] Mengambil ID dari URL: %s", idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[flow-4] Error: ID tidak valid - %v", err)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "ID tidak valid",
		})
		return
	}

	log.Printf("[flow-5] Menghapus kategori dengan ID: %d", id)
	success := store.DeleteKategori(id)

	if !success {
		log.Printf("[flow-6] Kategori dengan ID %d tidak ditemukan", id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error": fmt.Sprintf("Kategori dengan ID %d tidak ditemukan", id),
		})
		return
	}

	log.Println("[flow-7] Kategori berhasil dihapus")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Kategori berhasil dihapus",
	})
}
