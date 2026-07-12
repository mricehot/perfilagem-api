package service

import (
	"errors"
	"perfilagem-api/internal/models"
	"perfilagem-api/internal/store"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrNomeObrigatorio   = errors.New("O nome do anel é obrigatório")
	ErrNomeDuplicado     = errors.New("Já existe um anel com esse nome")
	ErrNomeNaoEncontrado = errors.New("Anel não encontrado")
)

type AnelService struct {
	store *store.AnelStore
}

func NewAnelService(store *store.AnelStore) *AnelService {
	return &AnelService{store: store}
}

func (s *AnelService) Criar(anel models.Anel) (models.Anel, error) {
	if anel.Nome == "" {
		return models.Anel{}, ErrNomeObrigatorio
	}
	todos := s.store.ListarTodos()
	for _, existente := range todos {
		if strings.EqualFold(existente.Nome, anel.Nome) {
			return models.Anel{}, ErrNomeDuplicado
		}
	}

	anel.ID = uuid.NewString() // Função fictícia para gerar um ID único
	anel.Ativo = true
	s.store.Criar(anel)
	return anel, nil
}

func (s *AnelService) Remover(id string) bool {
	return s.store.Remover(id)
}

func (s *AnelService) ListarTodos() []models.Anel {
	return s.store.ListarTodos()
}

func (s *AnelService) Atualizar(id string, dadosNovos models.Anel) (models.Anel, error) {
	if dadosNovos.Nome == "" {
		return models.Anel{}, ErrNomeObrigatorio
	}

	// Verifica se o nome já existe em outro anel
	todos := s.store.ListarTodos()
	for _, existente := range todos {
		if strings.EqualFold(existente.Nome, dadosNovos.Nome) && existente.ID != id {
			return models.Anel{}, ErrNomeDuplicado
		}
	}

	anel, encontrado := s.store.AtualizarAtivo(id, dadosNovos)
	if !encontrado {
		return models.Anel{}, ErrNomeNaoEncontrado
	}
	return anel, nil
}

func (s *AnelService) BuscarPorID(id string) (models.Anel, bool) {
	return s.store.BuscarPorID(id)
}
