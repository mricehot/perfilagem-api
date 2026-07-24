package routes

import (
	"net/http"
	"perfilagem-api/internal/handler"
	"perfilagem-api/internal/service"
)

func RegistrarRotasFuros(mux *http.ServeMux, furoService *service.FuroService) {

	mux.HandleFunc("GET /furos/leque/{lequeID}", handler.HandlerListarFurosPorLeque(furoService))
	mux.HandleFunc("POST /furos", handler.HandlerCriarFuro(furoService))
	mux.HandleFunc("GET /furos/{id}", handler.HandlerBuscarFuro(furoService))
	mux.HandleFunc("PUT /furos/{id}", handler.HandlerAtualizarFuro(furoService))
	mux.HandleFunc("DELETE /furos/{id}", handler.HandlerDeletarFuro(furoService))
}
