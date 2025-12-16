package handlers

import (
	"encoding/json"
	"fmt"
	"go-rest/src/models"
	"net/http"
)

func BooleanHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"value": true}`)
}

func BooleanHandler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := models.Response{
		Value: true,
	}

	json.NewEncoder(w).Encode(response)
}
