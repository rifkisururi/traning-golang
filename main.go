package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Produk represents a product in the cashier system.
type Produk struct {
	ID    int    `json:"id"`    // ID unik untuk produk.
	Nama  string `json:"nama"`  // Nama produk yang tampil di API.
	Harga int    `json:"harga"` // Harga produk dalam satuan rupiah.
	Stok  int    `json:"stok"`  // Stok tersedia untuk produk ini.
}

// In-memory storage (temporary, replaced by a database later).
var produk = []Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
	{ID: 2, Nama: "Vit 1000ml", Harga: 3000, Stok: 40},
	{ID: 3, Nama: "kecap", Harga: 12000, Stok: 20},
}

// getProdukByID mengambil 1 produk berdasarkan ID dari path URL.
func getProdukByID(w http.ResponseWriter, r *http.Request) {
	// Ambil bagian ID dari path, lalu ubah ke integer.
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// Jika ID bukan angka, kirim error 400.
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for _, p := range produk {
		if p.ID == id {
			// Jika ketemu, balas dengan JSON produk tersebut.
			
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	// Jika tidak ketemu, balas error 404.
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// updateProduk mengganti data produk sesuai ID dari path URL.
func updateProduk(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari path dan ubah ke integer.
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// Decode JSON body ke struct Produk.
	var produkUpdate Produk
	err = json.NewDecoder(r.Body).Decode(&produkUpdate)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i := range produk {
		if produk[i].ID == id {
			// Pastikan ID tetap mengikuti URL (bukan body).
			produkUpdate.ID = id
			produk[i] = produkUpdate

			// Kirim hasil update sebagai JSON.
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produkUpdate)
			return
		}
	}

	// Jika tidak ketemu, balas error 404.
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// deleteProduk menghapus produk berdasarkan ID dari path URL.
func deleteProduk(w http.ResponseWriter, r *http.Request) {
	// Ambil ID dari path dan ubah ke integer.
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for i, p := range produk {
		if p.ID == id {
			// Hapus elemen dengan menggabungkan slice sebelum dan sesudah index.
			produk = append(produk[:i], produk[i+1:]...)

			// Kirim pesan sukses dalam JSON.
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}
	}

	// Jika tidak ketemu, balas error 404.
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

func main() {
	// Endpoint untuk operasi berdasarkan ID (GET/PUT/DELETE).
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getProdukByID(w, r)
		case http.MethodPut:
			updateProduk(w, r)
		case http.MethodDelete:
			deleteProduk(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Endpoint koleksi produk (GET semua, POST tambah).
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			// Balas semua data produk sebagai JSON.
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		case http.MethodPost:
			// Decode JSON body untuk produk baru.
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}

			// Buat ID baru secara sederhana (berdasarkan panjang slice).
			produkBaru.ID = len(produk) + 1
			produk = append(produk, produkBaru)

			// Balas data yang baru dibuat.
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(produkBaru)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Endpoint health check untuk memastikan server hidup.
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API Running",
		})
	})

	// Log sederhana saat server mulai jalan.
	fmt.Println("Server running di localhost:8080")

	// Jalankan HTTP server pada port 8080.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// Tampilkan error jika server gagal start.
		fmt.Println("gagal running server:", err)
	}
}
   