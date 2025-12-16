package handlers

import (
	"encoding/json"
	"fmt"
	"go-rest/src/models"
	"net/http"
)

func IntHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"value": 42}`)
}

func IntHandler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := models.Response{
		Value: 42,
	}

	json.NewEncoder(w).Encode(response)
}
