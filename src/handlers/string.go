package handlers

import (
	"encoding/json"
	"fmt"
	"go-rest/src/models"
	"net/http"
)

func StringHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "Olá, Go-Rest! Aqui é uma string."}`)
}

func StringHandler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := models.Response{
		Message: "Olá, Go-Rest! Aqui é uma string.",
	}

	json.NewEncoder(w).Encode(response)
}
