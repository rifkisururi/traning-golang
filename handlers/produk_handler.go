// Package handlers menyimpan HTTP handler untuk produk.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"kasir-api/models"
	"kasir-api/store"
)

// GetProdukByID menangani GET /api/produk/{id}.
func GetProdukByID(w http.ResponseWriter, r *http.Request) {
	// Log langkah alur data untuk request ini.
	log.Printf("[flow-1] GetProdukByID start method=%s path=%s", r.Method, r.URL.Path)

	// Ambil ID dari path URL dan ubah ke integer.
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	log.Printf("[flow-2] GetProdukByID parse id raw=%q", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[flow-3] GetProdukByID parse id failed err=%v", err)
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}
	log.Printf("[flow-3] GetProdukByID parsed id=%d", id)

	// Ambil data dari store dan kirim jika ditemukan.
	log.Printf("[flow-4] GetProdukByID call store.GetByID id=%d", id)
	p, ok := store.GetByID(id)
	if ok {
		log.Printf("[flow-5] GetProdukByID found id=%d", p.ID)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)
		return
	}

	log.Printf("[flow-5] GetProdukByID not found id=%d", id)
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// UpdateProduk menangani PUT /api/produk/{id}.
func UpdateProduk(w http.ResponseWriter, r *http.Request) {
	// Log langkah alur data untuk request ini.
	log.Printf("[flow-1] UpdateProduk start method=%s path=%s", r.Method, r.URL.Path)

	// Ambil ID dari path URL dan ubah ke integer.
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	log.Printf("[flow-2] UpdateProduk parse id raw=%q", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[flow-3] UpdateProduk parse id failed err=%v", err)
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}
	log.Printf("[flow-3] UpdateProduk parsed id=%d", id)

	// Decode JSON body ke struct Produk.
	var produkUpdate models.Produk
	log.Printf("[flow-4] UpdateProduk decode body")
	err = json.NewDecoder(r.Body).Decode(&produkUpdate)
	if err != nil {
		log.Printf("[flow-5] UpdateProduk decode failed err=%v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	log.Printf("[flow-5] UpdateProduk decoded nama=%s harga=%d stok=%d", produkUpdate.Nama, produkUpdate.Harga, produkUpdate.Stok)

	// Update data di store dan kirim hasilnya.
	log.Printf("[flow-6] UpdateProduk call store.Update id=%d", id)
	updated, ok := store.Update(id, produkUpdate)
	if ok {
		log.Printf("[flow-7] UpdateProduk updated id=%d", updated.ID)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updated)
		return
	}

	log.Printf("[flow-7] UpdateProduk not found id=%d", id)
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// DeleteProduk menangani DELETE /api/produk/{id}.
func DeleteProduk(w http.ResponseWriter, r *http.Request) {
	// Log langkah alur data untuk request ini.
	log.Printf("[flow-1] DeleteProduk start method=%s path=%s", r.Method, r.URL.Path)

	// Ambil ID dari path URL dan ubah ke integer.
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	log.Printf("[flow-2] DeleteProduk parse id raw=%q", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[flow-3] DeleteProduk parse id failed err=%v", err)
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}
	log.Printf("[flow-3] DeleteProduk parsed id=%d", id)

	// Hapus data di store lalu kirim status.
	log.Printf("[flow-4] DeleteProduk call store.Delete id=%d", id)
	ok := store.Delete(id)
	if ok {
		log.Printf("[flow-5] DeleteProduk deleted id=%d", id)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "sukses delete",
		})
		return
	}

	log.Printf("[flow-5] DeleteProduk not found id=%d", id)
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// ListProduk menangani GET /api/produk.
func ListProduk(w http.ResponseWriter, r *http.Request) {
	// Log langkah alur data untuk request ini.
	log.Printf("[flow-1] ListProduk start method=%s path=%s", r.Method, r.URL.Path)

	// Ambil seluruh data produk lalu kirim sebagai JSON.
	log.Printf("[flow-2] ListProduk call store.GetAll")
	data := store.GetAll()
	log.Printf("[flow-3] ListProduk total=%d", len(data))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// CreateProduk menangani POST /api/produk.
func CreateProduk(w http.ResponseWriter, r *http.Request) {
	// Log langkah alur data untuk request ini.
	log.Printf("[flow-1] CreateProduk start method=%s path=%s", r.Method, r.URL.Path)

	// Decode JSON body ke struct Produk.
	var produkBaru models.Produk
	log.Printf("[flow-2] CreateProduk decode body")
	err := json.NewDecoder(r.Body).Decode(&produkBaru)
	if err != nil {
		log.Printf("[flow-3] CreateProduk decode failed err=%v", err)
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	log.Printf("[flow-3] CreateProduk decoded nama=%s harga=%d stok=%d", produkBaru.Nama, produkBaru.Harga, produkBaru.Stok)

	// Buat ID baru dan simpan ke store.
	produkBaru.ID = store.NextID()
	log.Printf("[flow-4] CreateProduk next id=%d", produkBaru.ID)
	log.Printf("[flow-5] CreateProduk call store.Add id=%d", produkBaru.ID)
	created := store.Add(produkBaru)

	// Kirim data yang baru dibuat.
	log.Printf("[flow-6] CreateProduk created id=%d", created.ID)


	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}
