package store

import (
	"sync"

	"perfilagem-api/internal/models"
)

// AnelStore guarda os anéis em memória, protegido contra acesso concorrente
type AnelStore struct {
	mu    sync.Mutex
	dados map[string]models.Anel
}

// NewAnelStore cria um store vazio, pronto pra usar
func NewAnelStore() *AnelStore {
	return &AnelStore{
		dados: make(map[string]models.Anel),
	}
}

// Criar adiciona um novo anel ao store
func (s *AnelStore) Criar(anel models.Anel) {
	// 1. Trava o mutex antes de mexer no map:
	s.mu.Lock()
	// 2. Usa "
	defer s.mu.Unlock()
	// " logo em seguida (explico o "defer" abaixo)
	// 3. Salva no map:
	s.dados[anel.ID] = anel
}

// BuscarPorID retorna o anel e um bool dizendo se foi encontrado
func (s *AnelStore) BuscarPorID(id string) (models.Anel, bool) {
	// 1. Trava o mutex
	s.mu.Lock()
	// 2. defer Unlock
	defer s.mu.Unlock()
	// 3. Busca no map. Em Go, buscar em map retorna 2 valores:
	anel, encontrado := s.dados[id]
	//    "encontrado" é bool: true se a chave existe, false se não
	// 4. Retorna os dois: return anel, encontrado
	return anel, encontrado
}
