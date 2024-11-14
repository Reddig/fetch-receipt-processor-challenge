package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler handles the "/" route
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}