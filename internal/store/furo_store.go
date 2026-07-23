package store

import (
	"fmt"
	"perfilagem-api/internal/models"

	"gorm.io/gorm"
)

type FuroStore struct {
	db *gorm.DB
}

func NewFuroStore(db *gorm.DB) *FuroStore {
	return &FuroStore{
		db: db,
	}
}

func (s *FuroStore) Criar(furo models.Furo) error {
	err := s.db.Create(&furo).Error
	if err != nil {
		return fmt.Errorf("erro ao criar furo: %w", err)
	}
	return nil
}

// buscar por id
func (s *FuroStore) BuscarPorID(id string) (models.Furo, bool) {
	var furo models.Furo
	resultado := s.db.First(&furo, "id = ?", id)
	if resultado.Error != nil {
		return models.Furo{}, false
	}
	return furo, true
}

// listar por leque
func (s *FuroStore) ListarPorLeque(lequeID string) ([]models.Furo, error) {
	var furos []models.Furo
	err := s.db.Where("leque_id = ?", lequeID).Find(&furos).Error
	if err != nil {
		return nil, fmt.Errorf("erro ao listar furos: %w", err)
	}
	return furos, nil
}

// atualizar furo
func (s *FuroStore) Atualizar(id string, dadosNovos models.Furo) (models.Furo, error) {
	var furoExistente models.Furo
	resultado := s.db.First(&furoExistente, "id = ?", id)
	if resultado.Error != nil {
		return models.Furo{}, fmt.Errorf("erro ao atualizar furo: %w", resultado.Error)
	}
	furoExistente.Numero = dadosNovos.Numero
	furoExistente.MetragemEsperada = dadosNovos.MetragemEsperada
	furoExistente.MetragemReal = dadosNovos.MetragemReal
	furoExistente.Situacao = dadosNovos.Situacao
	err := s.db.Save(&furoExistente).Error
	if err != nil {
		return models.Furo{}, fmt.Errorf("erro ao atualizar furo: %w", err)
	}
	return furoExistente, nil

}

// remover furo
func (s *FuroStore) Remover(id string) bool {
	var furo models.Furo
	resultado := s.db.First(&furo, "id = ?", id)
	if resultado.Error != nil {
		return false
	}
	err := s.db.Delete(&furo).Error
	if err != nil {
		return false
	}
	return true
}
