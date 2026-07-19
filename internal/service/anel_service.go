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
	ErrLequeAbertoNoAnel   = errors.New("Não é possível desativar o anel, pois existem leques abertos associados a ele")
)

type AnelService struct {
	anelStore  *store.AnelStore
	lequeStore *store.LequeStore
}

func NewAnelService(anelStore *store.AnelStore, lequeStore *store.LequeStore) *AnelService {
	return &AnelService{anelStore: anelStore, lequeStore: lequeStore}
}

func (s *AnelService) Criar(anel models.Anel) (models.Anel, error) {
	if anel.Nome == "" {
		return models.Anel{}, ErrNomeObrigatorio
	}
	todos, err := s.anelStore.ListarTodos()
	if err != nil {
		return models.Anel{}, err
	}

	for _, existente := range todos {
		if strings.EqualFold(existente.Nome, anel.Nome) {
			return models.Anel{}, ErrNomeDuplicado
		}
	}
	err = s.anelStore.DesativarTodos() // Desativa todos os anéis existentes antes de criar um novo
	if err != nil {
		return models.Anel{}, err
	}

	anel.ID = uuid.NewString() // Função fictícia para gerar um ID único
	anel.Ativo = true

	if err := s.anelStore.Criar(anel); err != nil {
		return models.Anel{}, err
	}

	return anel, nil
}

func (s *AnelService) ListarTodos() ([]models.Anel, error) {
	todos, err := s.anelStore.ListarTodos()
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
	todos, err := s.anelStore.ListarTodos()
	if err != nil {
		return models.Anel{}, err
	}
	for _, existente := range todos {
		if strings.EqualFold(existente.Nome, dadosNovos.Nome) && existente.ID != id {
			return models.Anel{}, ErrNomeDuplicado
		}
	}

	anelAtual, encontrado := s.anelStore.BuscarPorID(id)
	if !encontrado {
		return models.Anel{}, ErrNomeNaoEncontrado
	}

	if anelAtual.Ativo && !dadosNovos.Ativo {
		return models.Anel{}, ErrPrecisaTerAnelAtivo
	}

	if dadosNovos.Ativo {
		// Desativa todos os outros aneis antes de ativar o novo
		err := s.anelStore.DesativarTodos()
		if err != nil {
			return models.Anel{}, err
		}

	}

	anel, encontrado := s.anelStore.AtualizarAtivo(id, dadosNovos)
	if !encontrado {
		return models.Anel{}, ErrNomeNaoEncontrado
	}

	return anel, nil
}

func (s *AnelService) BuscarPorID(id string) (models.Anel, bool) {
	return s.anelStore.BuscarPorID(id)
}

func (s *AnelService) Remover(id string) error {
	anelRemovido, encontrado := s.anelStore.BuscarPorID(id)
	if !encontrado {
		return ErrNomeNaoEncontrado
	}

	// Verifica se existem leques abertos associados a esse anel
	leques, err := s.lequeStore.ListarPorAnel(id)
	if err != nil {
		return err
	}
	for _, leque := range leques {
		if leque.Status == "aberto" {
			return ErrLequeAbertoNoAnel
		}
	}

	// Remove o anel
	removido := s.anelStore.Remover(id)
	if anelRemovido.Ativo && removido {
		//ativa o proximo anel da lista, se houver
		todos, err := s.anelStore.ListarTodos()
		if err != nil {
			return err
		}
		if len(todos) == 0 {
			return nil // Nenhum anel restante para ativar
		}
		if len(todos) > 0 {
			// Ativa o primeiro anel da lista
			anelASerAtivado := todos[0]
			s.anelStore.AtualizarAtivo(anelASerAtivado.ID, models.Anel{Nome: anelASerAtivado.Nome, Ativo: true})
		}

	}

	if !removido {
		return ErrNomeNaoEncontrado
	}

	return nil
}
