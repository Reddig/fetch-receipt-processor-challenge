package main

import (
	"fmt"
	"log"
	"net/http"
	"fetch-receipt-processor-challenge/src/routes" // Import the routes package
)

func main() {
	// Initialize the router from the routes package
	router := routes.InitRouter()

	// Start the server
	fmt.Println("Server listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}