package routes

import (
	"fetch-receipt-processor-challenge/src/handlers"

	"github.com/gorilla/mux"
)

// InitRouter initializes the Mux router and routes
func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// Map routes to handlers
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/receipts/{id}/points", handlers.ReceiptPointsHandler).Methods("GET")
	router.HandleFunc("/receipts/{id}", handlers.ReceiptHandler).Methods("GET")
	router.HandleFunc("/receipts", handlers.ReceiptsHandler).Methods("GET")
	router.HandleFunc("/receipts/process", handlers.AddReceiptHandler).Methods("POST")

	return router
}
