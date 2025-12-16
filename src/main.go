package main

import (
	"fmt"
	"go-rest/src/handlers"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "----- Go Rest | Main -----")
}

func main() {
	http.HandleFunc("/api/old", index)
	//http.ListenAndServe(":8080", nil)

	http.HandleFunc("/api", handlers.IndexHandler)
	http.HandleFunc("/api/string", handlers.StringHandler2)
	http.HandleFunc("/api/int", handlers.IntHandler2)
	http.HandleFunc("/api/float", handlers.FloatHandler2)
	http.HandleFunc("/api/boolean", handlers.BooleanHandler2)
	http.HandleFunc("/api/user", handlers.GetUser)
	http.HandleFunc("/api/users", handlers.GetUsers)

	log.Println("Servidor rodando na porta 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
