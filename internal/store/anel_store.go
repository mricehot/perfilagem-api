package store

import (
	"fmt"
	"perfilagem-api/internal/models"

	"gorm.io/gorm"
)

// AnelStore guarda os anéis em memória, protegido contra acesso concorrente
type AnelStore struct {
	db *gorm.DB
}

func NewAnelStore(db *gorm.DB) *AnelStore {
	return &AnelStore{
		db: db,
	}
}

// Criar adiciona um novo anel ao store
func (s *AnelStore) Criar(anel models.Anel) error {
	err := s.db.Create(&anel).Error
	if err != nil {
		return fmt.Errorf("erro ao criar anel: %w", err)
	}
	return nil
}
func (s *AnelStore) BuscarPorID(id string) (models.Anel, bool) {
	var anel models.Anel

	resultado := s.db.First(&anel, "id = ?", id)
	if resultado.Error != nil {
		return models.Anel{}, false
	}

	return anel, true
}
func (s *AnelStore) ListarTodos() ([]models.Anel, error) {
	var aneis []models.Anel

	// dica: s.db.Find(&aneis) preenche o slice inteiro com todos os registros da tabela
	// devolve um *gorm.DB, e você usa .Error pra saber se deu problema
	err := s.db.Find(&aneis).Error
	if err != nil {
		return nil, fmt.Errorf("erro ao listar aneis: %w", err)
	}
	return aneis, nil
}

// tenta escrever a chamada e o tratamento do erro

// Atualizar substitui os dados de um anel existente. Retorna false se o ID não existir.
func (s *AnelStore) AtualizarAtivo(id string, dadosNovos models.Anel) (models.Anel, bool) {
	var anel models.Anel
	resultado := s.db.First(&anel, "id = ?", id)
	if resultado.Error != nil {
		return models.Anel{}, false
	}

	// Atualiza os campos do anel existente com os novos dados
	anel.Nome = dadosNovos.Nome
	anel.Ativo = dadosNovos.Ativo
	err := s.db.Save(&anel).Error
	if err != nil {
		return models.Anel{}, false
	}
	return anel, true
}

func (s *AnelStore) Remover(id string) bool {
	var anel models.Anel
	resultado := s.db.First(&anel, "id = ?", id)
	if resultado.Error != nil {
		return false
	}
	err := s.db.Delete(&anel).Error
	if err != nil {
		return false
	}
	return true
}

func (s *AnelStore) DesativarTodos() error {

	err := s.db.Model(&models.Anel{}).Where("ativo = ?", true).Update("ativo", false).Error
	if err != nil {
		return fmt.Errorf("erro ao desativar todos os anéis: %w", err)
	}

	return nil
}
