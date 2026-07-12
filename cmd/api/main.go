package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"perfilagem-api/internal/models"
	"perfilagem-api/internal/service"
	"perfilagem-api/internal/store"
)

func main() {
	mux := http.NewServeMux()
	anelStore := store.NewAnelStore()
	anelService := service.NewAnelService(anelStore)

	mux.HandleFunc("/health", handlerHealth)
	mux.HandleFunc("GET /aneis/{id}", handlerBuscarAnel(anelService))
	mux.HandleFunc("POST /aneis", handlerCriarAnel(anelService))
	mux.HandleFunc("GET /aneis", handlerListarAneis(anelService))
	mux.HandleFunc("PUT /aneis/{id}", handlerAtualizarAnel(anelService))
	mux.HandleFunc("DELETE /aneis/{id}", handlerDeletarAnel(anelService))

	log.Println("Servidor rodando na porta 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")

}

func handlerBuscarAnel(s *service.AnelService) http.HandlerFunc {
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
func handlerCriarAnel(s *service.AnelService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var anel models.Anel

		err := json.NewDecoder(r.Body).Decode(&anel)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		anel, err = s.Criar(anel)
		if err != nil {
			if err == service.ErrNomeObrigatorio || err == service.ErrNomeDuplicado {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, "Erro interno", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(anel)

	}
}

func handlerListarAneis(s *service.AnelService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Chama s.ListarTodos() e guarda o resultado numa variável (ex: "aneis")
		aneis := s.ListarTodos()
		// 2. Content-Type json (você já sabe fazer isso)

		// 3. Encode do slice "aneis" na resposta
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(aneis)
	}
}

func handlerAtualizarAnel(s *service.AnelService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. pega o id da URL (você já sabe fazer isso)
		id := r.PathValue("id")
		// 2. decodifica o corpo JSON pra uma variável models.Anel (você já sabe fazer isso)
		var dadosNovos models.Anel
		err := json.NewDecoder(r.Body).Decode(&dadosNovos)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		// 3. chama s.AtualizarAtivo(id, dadosNovos) - guarda o bool retornado
		anelAtualizado, err := s.Atualizar(id, dadosNovos)
		if err != nil {
			if err == service.ErrNomeNaoEncontrado {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if err == service.ErrNomeObrigatorio || err == service.ErrNomeDuplicado {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, "Erro interno", http.StatusInternalServerError)
			return
		}

		// 5. se existir, responde 200 (padrão, não precisa WriteHeader) com o anel atualizado
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(anelAtualizado)
	}
}
func handlerDeletarAnel(s *service.AnelService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		deletado := s.Remover(id)
		if !deletado {
			http.Error(w, "Anel não encontrado", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
