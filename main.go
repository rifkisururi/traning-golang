// Package main adalah entry point untuk server HTTP.
package main

import (
	"log"
	"net/http"

	"kasir-api/handlers"
)

// main mendaftarkan route dan menyalakan server.
func main() {
	// Atur format log agar urutan waktu terlihat jelas.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// Endpoint untuk operasi berdasarkan ID (GET/PUT/DELETE).
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetProdukByID(w, r)
		case http.MethodPut:
			handlers.UpdateProduk(w, r)
		case http.MethodDelete:
			handlers.DeleteProduk(w, r)
		default:
			log.Printf("[flow-0] Method not allowed method=%s path=%s", r.Method, r.URL.Path)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Endpoint koleksi produk (GET semua, POST tambah).
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.ListProduk(w, r)
		case http.MethodPost:
			handlers.CreateProduk(w, r)
		default:
			log.Printf("[flow-0] Method not allowed method=%s path=%s", r.Method, r.URL.Path)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Endpoint health check untuk memastikan server hidup.
	http.HandleFunc("/health", handlers.Health)

	// Log sederhana saat server mulai jalan.
	log.Printf("[flow-0] Server running di localhost:8080")

	// Jalankan HTTP server pada port 8080.
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// Tampilkan error jika server gagal start.
		log.Printf("[flow-0] gagal running server: %v", err)
	}
}
