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

// ListarTodos devolve todos os anéis guardados, sem ordem garantida
func (s *AnelStore) ListarTodos() []models.Anel {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 1. Declara um slice vazio: var lista []models.Anel
	var lista []models.Anel

	// 2. Usa "range" pra percorrer s.dados (que é um map[string]models.Anel)
	//    Lembre: range num map devolve (chave, valor) a cada volta.
	for _, anel := range s.dados {
		lista = append(lista, anel)
	}
	//    Você só precisa do valor aqui, pode ignorar a chave usando "_"
	//    for _, anel := range s.dados { ... }

	// 3. Dentro do for, adiciona cada anel na lista: lista = append(lista, anel)

	// 4. Depois do for, retorna a lista: return lista
	return lista
}

// Atualizar substitui os dados de um anel existente. Retorna false se o ID não existir.
func (s *AnelStore) AtualizarAtivo(id string, dadosNovos models.Anel) (models.Anel, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	// pensa: antes de sobrescrever, você precisa checar se esse ID já existe no map.
	_, encontrado := s.dados[id]
	if !encontrado {
		return models.Anel{}, false
	}

	// se não existir, devolve false sem alterar nada.
	// se existir, você precisa manter o ID original (dadosNovos pode não ter vindo com ID preenchido),
	// sobrescrever no map, e devolver true.
	dadosNovos.ID = id
	s.dados[id] = dadosNovos
	return dadosNovos, true
}

func (s *AnelStore) Remover(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, encontrado := s.dados[id]
	if !encontrado {
		return false
	}
	delete(s.dados, id)
	return true
}
