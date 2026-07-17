package routes

import (
	"net/http"
	"perfilagem-api/internal/service"
)

type Services struct {
	AnelService  *service.AnelService
	LequeService *service.LequeService
}

func NovoRouter(svcs *Services) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	RegistrarRotasAnel(mux, svcs.AnelService)
	RegistrarRotasLeque(mux, svcs.LequeService)
	return mux
}
