package handlers

import (
	"encoding/json"
	"go-rest/src/models"
	"net/http"
	"strconv"
	"time"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	// Recuperando o ID da URL
	idParam := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	// Simulando um usuário (em um banco de dados real, você faria uma busca)
	user := models.User{
		ID:           id,
		Name:         "User " + strconv.Itoa(id),
		CreationDate: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Simulando a lista de usuários
	users := []models.User{
		{ID: 1, Name: "User 1", CreationDate: time.Now().Add(-24 * time.Hour)},
		{ID: 2, Name: "User 2", CreationDate: time.Now().Add(-48 * time.Hour)},
		{ID: 3, Name: "User 3", CreationDate: time.Now().Add(-72 * time.Hour)},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
