package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"perfilagem-api/internal/models"
	"perfilagem-api/internal/store"
)

func main() {
	mux := http.NewServeMux()
	anelStore := store.NewAnelStore()

	mux.HandleFunc("/health", handlerHealth)
	mux.HandleFunc("GET /aneis/{id}", handlerBuscarAnel(anelStore))
	mux.HandleFunc("POST /aneis", handlerCriarAnel(anelStore))

	log.Println("Servidor rodando na porta 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")

}

func handlerBuscarAnel(s *store.AnelStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		anel, encontrado := s.BuscarPorID(id)
		if !encontrado {
			http.Error(w, "Anel não encontrado", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(anel)
	}
}
func handlerCriarAnel(store *store.AnelStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var anel models.Anel

		err := json.NewDecoder(r.Body).Decode(&anel)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		anel.ID = "fake-id-123"
		anel.Ativo = true
		store.Criar(anel)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(anel)

	}
}
