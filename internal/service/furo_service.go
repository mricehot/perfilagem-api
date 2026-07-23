package service

import (
	"errors"
	"fmt"
	"perfilagem-api/internal/models"
	"perfilagem-api/internal/store"

	"github.com/google/uuid"
)

var (
	ErrFuroNaoEncontrado               = errors.New("Furo não encontrado")
	ErrLequeNaoEncontrado              = errors.New("Leque não encontrado")
	ErrLequeFechado                    = errors.New("Não é possível criar furos em leques fechados")
	ErrFuroDuplicado                   = errors.New("Já existe um furo com esse número no leque")
	ErrFuroNumeroObrigatorio           = errors.New("O número do furo é obrigatório")
	ErrFuroMetragemEsperadaObrigatoria = errors.New("A metragem esperada do furo é obrigatória")
)

type FuroService struct {
	furoStore  *store.FuroStore
	lequeStore *store.LequeStore
}

func NewFuroService(furoStore *store.FuroStore, lequeStore *store.LequeStore) *FuroService {
	return &FuroService{
		furoStore:  furoStore,
		lequeStore: lequeStore,
	}
}

func (s *FuroService) Criar(furo models.Furo) (models.Furo, error) {
	if furo.Numero == "" {
		return models.Furo{}, ErrFuroNumeroObrigatorio
	}

	if furo.MetragemEsperada == 0 {
		return models.Furo{}, ErrFuroMetragemEsperadaObrigatoria
	}

	leque, encontrado := s.lequeStore.BuscarPorID(furo.LequeID)
	if !encontrado {
		return models.Furo{}, ErrLequeNaoEncontrado
	}
	if leque.Status == "fechado" {
		return models.Furo{}, ErrLequeFechado
	}

	furosExistentes, err := s.furoStore.ListarPorLeque(furo.LequeID)
	if err != nil {
		return models.Furo{}, err
	}
	for _, existente := range furosExistentes {
		if existente.Numero == furo.Numero {
			return models.Furo{}, ErrFuroDuplicado
		}
	}

	furo.ID = uuid.NewString() // Função fictícia para gerar um ID único

	if furo.Situacao == "" {
		furo.Situacao = "livre"
	}
	err = s.furoStore.Criar(furo)
	if err != nil {
		return models.Furo{}, err
	}

	return furo, nil
}

func (s *FuroService) BuscarPorID(id string) (models.Furo, bool) {

	furo, encontrado := s.furoStore.BuscarPorID(id)

	if !encontrado {
		return models.Furo{}, false
	}

	return furo, true
}

func (s *FuroService) ListarPorLeque(lequeID string) ([]models.Furo, error) {

	furos, err := s.furoStore.ListarPorLeque(lequeID)
	if err != nil {
		return nil, fmt.Errorf("erro ao listar furos: %w", err)
	}

	return furos, nil
}

func (s *FuroService) Atualizar(id string, dadosNovos models.Furo) (models.Furo, error) {
	furo, encontrado := s.furoStore.BuscarPorID(id)
	if !encontrado {
		return models.Furo{}, ErrFuroNaoEncontrado
	}
	leque, encontrado := s.lequeStore.BuscarPorID(furo.LequeID)
	if !encontrado {
		return models.Furo{}, ErrLequeNaoEncontrado
	}
	if leque.Status == "fechado" {
		return models.Furo{}, ErrLequeFechado
	}
	if dadosNovos.Numero == "" {
		return models.Furo{}, ErrFuroNumeroObrigatorio
	}
	if dadosNovos.MetragemEsperada == 0 {
		return models.Furo{}, ErrFuroMetragemEsperadaObrigatoria
	}
	if dadosNovos.Numero != furo.Numero {
		furosExistentes, err := s.furoStore.ListarPorLeque(furo.LequeID)
		if err != nil {
			return models.Furo{}, err
		}
		for _, existente := range furosExistentes {
			if existente.Numero == dadosNovos.Numero {
				return models.Furo{}, ErrFuroDuplicado
			}
		}
	}
	if dadosNovos.Situacao != "" {
		furo.Situacao = dadosNovos.Situacao
	}
	furo.Numero = dadosNovos.Numero
	furo.MetragemEsperada = dadosNovos.MetragemEsperada
	furo.MetragemReal = dadosNovos.MetragemReal

	furoAtualizado, err := s.furoStore.Atualizar(id, furo)
	if err != nil {
		return models.Furo{}, err
	}

	return furoAtualizado, nil

}

func (s *FuroService) Remover(id string) error {
	furo, encontrado := s.furoStore.BuscarPorID(id)
	if !encontrado {
		return ErrFuroNaoEncontrado
	}
	encontradoLeque, encontrado := s.lequeStore.BuscarPorID(furo.LequeID)
	if !encontrado {
		return ErrLequeNaoEncontrado
	}
	if encontradoLeque.Status == "fechado" {
		return ErrLequeFechado
	}

	removido := s.furoStore.Remover(id)
	if !removido {
		return ErrFuroNaoEncontrado
	}

	return nil
}
