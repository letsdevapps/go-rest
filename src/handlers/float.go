package handlers

import (
	"encoding/json"
	"fmt"
	"go-rest/src/models"
	"net/http"
)

func FloatHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"value": 3.14}`)
}

func FloatHandler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := models.Response{
		Value: 3.14,
	}

	json.NewEncoder(w).Encode(response)
}
