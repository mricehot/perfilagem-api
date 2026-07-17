package routes

import (
	"net/http"

	"perfilagem-api/internal/handler"
	"perfilagem-api/internal/service"
)

func RegistrarRotasAnel(mux *http.ServeMux, s *service.AnelService) {
	mux.HandleFunc("GET /aneis/{id}", handler.HandlerBuscarAnel(s))
	mux.HandleFunc("POST /aneis", handler.HandlerCriarAnel(s))
	mux.HandleFunc("GET /aneis", handler.HandlerListarAneis(s))
	mux.HandleFunc("PUT /aneis/{id}", handler.HandlerAtualizarAnel(s))
	mux.HandleFunc("DELETE /aneis/{id}", handler.HandlerDeletarAnel(s))
}
