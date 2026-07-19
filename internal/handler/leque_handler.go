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
		id := r.PathValue("id")
		removido := s.Remover(id)
		if !removido {
			http.Error(w, "Leque não encontrado", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}

}

func HandlerFinalizarLeque(s *service.LequeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		leque, err := s.Finalizar(id)
		if err != nil {
			if err == service.ErrNumeroNaoEncontrado {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			http.Error(w, "Erro interno", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(leque)
	}
}

func HandlerReabrirLeque(s *service.LequeService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		leque, err := s.Reabrir(id)
		if err != nil {
			if err == service.ErrNumeroNaoEncontrado || err == service.ErrAnelNaoEncontrado {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if err == service.ErrLequeAbertoExiste {
				http.Error(w, err.Error(), http.StatusConflict)
				return
			}
			http.Error(w, "Erro interno", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(leque)
	}
}
