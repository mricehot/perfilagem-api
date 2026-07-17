package handler

import (
	"encoding/json"
	"net/http"
	"perfilagem-api/internal/models"
	"perfilagem-api/internal/service"
)

func HandlerBuscarLeque(s *service.LequeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementação do handler para buscar um leque
		id := r.PathValue("id")
		leque, err := s.BuscarPorID(id)
		if err != nil {
			http.Error(w, "Leque não encontrado", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(leque)

	}
}

func HandlerCriarLeque(s *service.LequeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementação do handler para criar um leque
		var leque models.Leque
		err := json.NewDecoder(r.Body).Decode(&leque)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		leque, err = s.Criar(leque)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(leque)

	}

}

func HandlerListarLequesPorAnel(s *service.LequeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementação do handler para listar leques
		anelID := r.PathValue("anelID")
		leques, err := s.ListarPorAnel(anelID)
		if err != nil {
			http.Error(w, "Erro interno", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(leques)
	}

}

func HandlerAtualizarLeque(s *service.LequeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementação do handler para atualizar um leque
		id := r.PathValue("id")
		var leque models.Leque
		err := json.NewDecoder(r.Body).Decode(&leque)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		leque, err = s.Atualizar(id, leque)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(leque)
	}

}

func HandlerDeletarLeque(s *service.LequeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Implementação do handler para deletar um leque
	}

}
