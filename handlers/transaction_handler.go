// Package handlers menyimpan HTTP handler untuk transaksi.
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

// HandleCheckout menangani /api/checkout (POST).
func HandleCheckout(w http.ResponseWriter, r *http.Request) {
	log.Printf("[flow-1] HandleCheckout start method=%s path=%s", r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodPost:
		Checkout(w, r)
	default:
		log.Printf("[flow-2] HandleCheckout method not allowed method=%s", r.Method)
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Checkout menangani POST /api/checkout.
func Checkout(w http.ResponseWriter, r *http.Request) {
	log.Printf("[flow-1] Checkout start method=%s path=%s", r.Method, r.URL.Path)

	// Decode request body.
	var req models.CheckoutRequest
	log.Printf("[flow-2] Checkout decode body")
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Printf("[flow-3] Checkout decode failed err=%v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validasi request.
	if len(req.Items) == 0 {
		log.Printf("[flow-3] Checkout empty items")
		http.Error(w, "Items cannot be empty", http.StatusBadRequest)
		return
	}

	log.Printf("[flow-3] Checkout items count=%d", len(req.Items))

	// Panggil store untuk membuat transaksi.
	transaction, err := store.CreateTransaction(req.Items)
	if err != nil {
		log.Printf("[flow-4] Checkout create transaction failed err=%v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("[flow-5] Checkout success id=%d total=%d", transaction.ID, transaction.TotalAmount)

	// Kirim response.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(transaction)
}

// GetTransactionByID menangani GET /api/transaction/{id}.
func GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[flow-1] GetTransactionByID start method=%s path=%s", r.Method, r.URL.Path)

	// Ambil ID dari path URL.
	idStr := strings.TrimPrefix(r.URL.Path, "/api/transaction/")
	log.Printf("[flow-2] GetTransactionByID parse id raw=%q", idStr)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[flow-3] GetTransactionByID parse id failed err=%v", err)
		http.Error(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}
	log.Printf("[flow-3] GetTransactionByID parsed id=%d", id)

	// Ambil data dari store.
	transaction, err := store.GetTransactionByID(id)
	if err != nil {
		log.Printf("[flow-4] GetTransactionByID not found id=%d err=%v", id, err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	log.Printf("[flow-5] GetTransactionByID found id=%d total=%d", transaction.ID, transaction.TotalAmount)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}

// GetAllTransactions menangani GET /api/transaction.
func GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	log.Printf("[flow-1] GetAllTransactions start method=%s path=%s", r.Method, r.URL.Path)

	transactions, err := store.GetAllTransactions()
	if err != nil {
		log.Printf("[flow-2] GetAllTransactions failed err=%v", err)
		http.Error(w, "Failed to get transactions", http.StatusInternalServerError)
		return
	}

	log.Printf("[flow-3] GetAllTransactions total=%d", len(transactions))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transactions)
}
