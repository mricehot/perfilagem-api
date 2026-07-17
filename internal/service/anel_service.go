package service

import (
	"errors"
	"perfilagem-api/internal/models"
	"perfilagem-api/internal/store"
	"strings"

	"github.com/google/uuid"
)

var (
	ErrNomeObrigatorio     = errors.New("O nome do anel é obrigatório")
	ErrNomeDuplicado       = errors.New("Já existe um anel com esse nome")
	ErrNomeNaoEncontrado   = errors.New("Anel não encontrado")
	ErrPrecisaTerAnelAtivo = errors.New("não é possível desativar o único anel ativo")
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
	todos, err := s.store.ListarTodos()
	if err != nil {
		return models.Anel{}, err
	}

	for _, existente := range todos {
		if strings.EqualFold(existente.Nome, anel.Nome) {
			return models.Anel{}, ErrNomeDuplicado
		}
	}
	err = s.store.DesativarTodos() // Desativa todos os anéis existentes antes de criar um novo
	if err != nil {
		return models.Anel{}, err
	}

	anel.ID = uuid.NewString() // Função fictícia para gerar um ID único
	anel.Ativo = true

	if err := s.store.Criar(anel); err != nil {
		return models.Anel{}, err
	}

	return anel, nil
}

func (s *AnelService) Remover(id string) bool {
	return s.store.Remover(id)
}

func (s *AnelService) ListarTodos() ([]models.Anel, error) {
	todos, err := s.store.ListarTodos()
	if err != nil {
		return []models.Anel{}, err
	}
	return todos, nil
}

func (s *AnelService) Atualizar(id string, dadosNovos models.Anel) (models.Anel, error) {
	if dadosNovos.Nome == "" {
		return models.Anel{}, ErrNomeObrigatorio
	}

	// Verifica se o nome já existe em outro anel
	todos, err := s.store.ListarTodos()
	if err != nil {
		return models.Anel{}, err
	}
	for _, existente := range todos {
		if strings.EqualFold(existente.Nome, dadosNovos.Nome) && existente.ID != id {
			return models.Anel{}, ErrNomeDuplicado
		}
	}

	anelAtual, encontrado := s.store.BuscarPorID(id)
	if !encontrado {
		return models.Anel{}, ErrNomeNaoEncontrado
	}

	if anelAtual.Ativo && !dadosNovos.Ativo {
		return models.Anel{}, ErrPrecisaTerAnelAtivo
	}

	if dadosNovos.Ativo {
		// Desativa todos os outros aneis antes de ativar o novo
		err := s.store.DesativarTodos()
		if err != nil {
			return models.Anel{}, err
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
