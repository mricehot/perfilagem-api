package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"perfilagem-api/internal/models"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handlerHealth)
	mux.HandleFunc("GET /aneis/{id}", handlerBuscarAnel)

	log.Println("Servidor rodando na porta 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")

}

func handlerBuscarAnel(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	anel := models.Anel{
		ID:    id,
		Nome:  "Galeria Teste",
		Ativo: true,
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(anel)
}
