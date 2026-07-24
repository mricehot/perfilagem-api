package handler

import (
	"encoding/json"
	"net/http"
	"perfilagem-api/internal/models"
	"perfilagem-api/internal/service"
)

func HandlerBuscarFuro(s *service.FuroService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		furo, encontrado := s.BuscarPorID(id)
		if !encontrado {
			http.Error(w, "Furo não encontrado", http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(furo)
	}
}

func HandlerCriarFuro(s *service.FuroService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var furo models.Furo

		err := json.NewDecoder(r.Body).Decode(&furo)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		furo, err = s.Criar(furo)
		if err != nil {
			if err == service.ErrLequeNaoEncontrado {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if err == service.ErrLequeFechado {
				http.Error(w, err.Error(), http.StatusConflict)
				return
			}
			if err == service.ErrFuroDuplicado || err == service.ErrFuroNumeroObrigatorio || err == service.ErrFuroMetragemEsperadaObrigatoria {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			http.Error(w, "Erro interno", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(furo)

	}
}

func HandlerListarFurosPorLeque(s *service.FuroService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lequeID := r.PathValue("lequeID")
		furos, err := s.ListarPorLeque(lequeID)
		if err != nil {
			http.Error(w, "Erro ao listar furos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(furos)
	}
}

func HandlerDeletarFuro(s *service.FuroService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		err := s.Remover(id)
		if err != nil {
			if err == service.ErrFuroNaoEncontrado || err == service.ErrLequeNaoEncontrado {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if err == service.ErrLequeFechado {
				http.Error(w, err.Error(), http.StatusConflict)
				return
			}

			http.Error(w, "Erro ao deletar furo", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func HandlerAtualizarFuro(s *service.FuroService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		var dadosNovos models.Furo

		err := json.NewDecoder(r.Body).Decode(&dadosNovos)
		if err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}

		furoAtualizado, err := s.Atualizar(id, dadosNovos)
		if err != nil {
			if err == service.ErrFuroNaoEncontrado || err == service.ErrLequeNaoEncontrado {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			if err == service.ErrLequeFechado {
				http.Error(w, err.Error(), http.StatusConflict)
				return
			}
			if err == service.ErrFuroDuplicado || err == service.ErrFuroNumeroObrigatorio || err == service.ErrFuroMetragemEsperadaObrigatoria {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			http.Error(w, "Erro ao atualizar furo", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(furoAtualizado)
	}
}
