package handlers

import (
	"encoding/json"
	"fetch-receipt-processor-challenge/src/models"
	"net/http"
	"github.com/gorilla/mux"
)

// ReceiptsHandler lists all receipts
func ReceiptsHandler(w http.ResponseWriter, r *http.Request) {
	receipts := models.GetAllReceipts()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receipts)
}

// ReceiptHandler gets a single receipt by ID
func ReceiptHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	receiptID := vars["id"]

	receipt, exists := models.GetReceipt(receiptID)
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receipt)
}

// ReceiptHandler gets a single receipt by ID
func ReceiptPointsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	receiptID := vars["id"]

	receipt, exists := models.GetReceipt(receiptID)
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	var points = models.CalculatePoints(receipt)
	response := map[string]int{
		"points": points,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// AddReceiptHandler adds a new receipt to in-memory storage
func AddReceiptHandler(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		json.NewEncoder(w).Encode(err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := models.ValidateReceipt(receipt); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	id, err := models.AddReceipt(receipt)
	if err != nil { // this is repeat code that I would refactor into a new function but I am learning Go as I am doing this
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}
	response := map[string]string{
		"id": id,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
