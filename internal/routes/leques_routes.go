package routes

import (
	"net/http"
	"perfilagem-api/internal/handler"
	"perfilagem-api/internal/service"
)

func RegistrarRotasLeque(mux *http.ServeMux, s *service.LequeService) {
	mux.HandleFunc("GET /leques/{id}", handler.HandlerBuscarLeque(s))
	mux.HandleFunc("POST /leques", handler.HandlerCriarLeque(s))
	mux.HandleFunc("GET /leques/anel/{anelID}", handler.HandlerListarLequesPorAnel(s))
	mux.HandleFunc("PUT /leques/{id}", handler.HandlerAtualizarLeque(s))
	mux.HandleFunc("DELETE /leques/{id}", handler.HandlerDeletarLeque(s))
}
