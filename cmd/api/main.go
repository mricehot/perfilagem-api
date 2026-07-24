package main

import (
	"log"
	"net/http"

	"perfilagem-api/internal/db"
	"perfilagem-api/internal/routes"
	"perfilagem-api/internal/service"
	"perfilagem-api/internal/store"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	gormDB, err := db.Conectar()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Conexão com o banco de dados estabelecida com sucesso!")

	anelStore := store.NewAnelStore(gormDB)
	lequeStore := store.NewLequeStore(gormDB)
	furoStore := store.NewFuroStore(gormDB)
	anelService := service.NewAnelService(anelStore, lequeStore)
	lequeService := service.NewLequeService(lequeStore, anelStore)
	furoService := service.NewFuroService(furoStore, lequeStore)

	services := routes.Services(routes.Services{
		AnelService:  anelService,
		LequeService: lequeService,
		FuroService:  furoService,
	})
	mux := routes.NovoRouter(&services)

	log.Println("Servidor rodando na porta 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
