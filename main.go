// Package main adalah entry point untuk server HTTP.
package main

import (
	"log"
	"net/http"

	"kasir-api/database"
	"kasir-api/handlers"
)

// main mendaftarkan route dan menyalakan server.
func main() {
	// Atur format log agar urutan waktu terlihat jelas.
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	// Load konfigurasi dari environment variables atau file .env
	config := LoadConfig()

	// Koneksi ke database
	err := database.ConnectDatabase(config.GetDBConnectionString())
	if err != nil {
		log.Fatalf("[main] Gagal koneksi ke database: %v", err)
	}
	defer database.CloseDatabase()

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

	// Endpoint untuk operasi kategori berdasarkan ID (GET/PUT/DELETE).
	http.HandleFunc("/api/kategori/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetKategoriByIDHandler(w, r)
		case http.MethodPut:
			handlers.UpdateKategoriHandler(w, r)
		case http.MethodDelete:
			handlers.DeleteKategoriHandler(w, r)
		default:
			log.Printf("[flow-0] Method not allowed method=%s path=%s", r.Method, r.URL.Path)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Endpoint koleksi kategori (GET semua, POST tambah).
	http.HandleFunc("/api/kategori", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetKategoriHandler(w, r)
		case http.MethodPost:
			handlers.CreateKategoriHandler(w, r)
		default:
			log.Printf("[flow-0] Method not allowed method=%s path=%s", r.Method, r.URL.Path)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Endpoint health check untuk memastikan server hidup.
	http.HandleFunc("/health", handlers.Health)

	// Redirect /swagger ke /swagger/ agar path relatif di UI bekerja.
	http.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[flow-0] Swagger redirect method=%s path=%s", r.Method, r.URL.Path)
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})

	// Static file server untuk Swagger UI dan OpenAPI spec.
	swaggerFiles := http.FileServer(http.Dir("swagger"))
	http.Handle("/swagger/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[flow-1] Swagger serve method=%s path=%s", r.Method, r.URL.Path)
		http.StripPrefix("/swagger/", swaggerFiles).ServeHTTP(w, r)
	}))

	// Log sederhana saat server mulai jalan.
	log.Printf("[flow-0] Server running di %s:%s", config.Host, config.Port)

	// Jalankan HTTP server dengan konfigurasi dari env atau .env
	addr := config.Host + ":" + config.Port
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		// Tampilkan error jika server gagal start.
		log.Printf("[flow-0] gagal running server: %v", err)
	}
}
