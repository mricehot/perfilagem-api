package service

import (
	"errors"
	"perfilagem-api/internal/models"
	"perfilagem-api/internal/store"

	"github.com/google/uuid"
)

var (
	ErrNumeroObrigatorio   = errors.New("O número do leque é obrigatório")
	ErrLequeDuplicado      = errors.New("Já existe um leque com esse número")
	ErrNumeroNaoEncontrado = errors.New("Leque não encontrado")
	ErrLequeAbertoExiste   = errors.New("Já existe um leque aberto para este anel")
	ErrAnelNaoEncontrado   = errors.New("Anel não encontrado")
)

type LequeService struct {
	store *store.LequeStore
	anel  *store.AnelStore
}

func NewLequeService(lequeStore *store.LequeStore, anelStore *store.AnelStore) *LequeService {
	return &LequeService{
		store: lequeStore,
		anel:  anelStore,
	}

}
func (s *LequeService) Criar(leque models.Leque) (models.Leque, error) {
	if leque.Numero == "" {
		return models.Leque{}, ErrNumeroObrigatorio
	}

	existentes, err := s.store.ListarPorAnel(leque.AnelID)
	if err != nil {
		return models.Leque{}, err
	}

	for _, existente := range existentes {
		if existente.Tipo == leque.Tipo && existente.Numero == leque.Numero {
			return models.Leque{}, ErrLequeDuplicado
		}
		if existente.Status == "aberto" {
			return models.Leque{}, ErrLequeAbertoExiste
		}
		// pensa: como você checa se "existente" está com status "aberto"?
		// se achar algum aberto, devolve ErrLequeAbertoExiste
	}

	leque.ID = uuid.NewString()
	leque.Status = "aberto"

	if err := s.store.Criar(leque); err != nil {
		return models.Leque{}, err
	}

	return leque, nil
}

func (s *LequeService) BuscarPorID(id string) (models.Leque, error) {

	leque, encontrado := s.store.BuscarPorID(id)
	if !encontrado {
		return models.Leque{}, ErrNumeroNaoEncontrado
	}

	// repassa direto pro store
	return leque, nil

}

func (s *LequeService) ListarPorAnel(anelID string) ([]models.Leque, error) {
	// repassa direto pro store
	return s.store.ListarPorAnel(anelID)
}

func (s *LequeService) Finalizar(id string) (models.Leque, error) {
	// chama s.store.AtualizarStatus(id, "fechado")
	_, encontrado := s.store.BuscarPorID(id)
	if !encontrado {
		return models.Leque{}, ErrNumeroNaoEncontrado
	}
	lequeAtualizado, _ := s.store.AtualizarStatus(id, "fechado")
	return lequeAtualizado, nil
	// se não encontrado, devolve ErrLequeNaoEncontrado
}

func (s *LequeService) Reabrir(id string) (models.Leque, error) {
	// aqui tem uma regra interessante: antes de reabrir esse leque,
	// você precisa checar se já existe outro leque ABERTO nesse mesmo anel.
	leque, encontrado := s.store.BuscarPorID(id)
	if !encontrado {
		return models.Leque{}, ErrNumeroNaoEncontrado
	}
	anelID := leque.AnelID

	_, anelEncontrado := s.anel.BuscarPorID(anelID)
	if !anelEncontrado {
		return models.Leque{}, ErrAnelNaoEncontrado
	}

	lequesDoAnel, err := s.store.ListarPorAnel(anelID)
	if err != nil {
		return models.Leque{}, err
	}
	for _, leque := range lequesDoAnel {
		if leque.Status == "aberto" && leque.ID != id {
			return models.Leque{}, ErrLequeAbertoExiste
		}
	}

	lequeAtualizado, _ := s.store.AtualizarStatus(id, "aberto")
	return lequeAtualizado, nil
	// se existir (e não for o mesmo "id"), o que fazer? bloquear, ou fechar o outro automaticamente?
	// por enquanto, pra simplificar, pode bloquear com ErrLequeAbertoExiste
	// depois disso, chama s.store.AtualizarStatus(id, "aberto")
}

func (s *LequeService) Remover(id string) bool {
	// repassa direto pro store
	return s.store.Remover(id)
}

func (s *LequeService) Atualizar(id string, leque models.Leque) (models.Leque, error) {
	_, encontrado := s.store.BuscarPorID(id)
	if !encontrado {
		return models.Leque{}, ErrNumeroNaoEncontrado
	}

	lequeAtualizado, err := s.store.Atualizar(id, leque)
	if err != nil {
		return models.Leque{}, err
	}

	return lequeAtualizado, nil
}
