// Package handlers menyimpan HTTP handler untuk health check.
package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

// Health menangani GET /health untuk status server.
func Health(w http.ResponseWriter, r *http.Request) {
	// Log langkah alur data untuk request ini.
	log.Printf("[flow-1] Health start method=%s path=%s", r.Method, r.URL.Path)

	// Set header JSON dan kirim status sederhana.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "OK",
		"message": "API Running",
	})
	log.Printf("[flow-2] Health respond status=OK")
}
