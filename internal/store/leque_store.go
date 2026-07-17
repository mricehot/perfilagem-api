package store

import (
	"fmt"
	"perfilagem-api/internal/models"

	"gorm.io/gorm"
)

type LequeStore struct {
	db *gorm.DB
}

func NewLequeStore(db *gorm.DB) *LequeStore {
	return &LequeStore{
		db: db,
	}
}

func (s *LequeStore) Criar(leque models.Leque) error {

	err := s.db.Create(&leque).Error
	if err != nil {
		return fmt.Errorf("erro ao criar leque: %w", err)
	}
	return nil
}

func (s *LequeStore) BuscarPorID(id string) (models.Leque, bool) {
	var leque models.Leque
	resultado := s.db.First(&leque, "id = ?", id)
	if resultado.Error != nil {
		return models.Leque{}, false
	}
	return leque, true
}

func (s *LequeStore) ListarPorAnel(id string) ([]models.Leque, error) {
	var leques []models.Leque
	err := s.db.Where("anel_id = ?", id).Find(&leques).Error
	if err != nil {
		return nil, fmt.Errorf("erro ao listar leques: %w", err)
	}
	return leques, nil
}

func (s *LequeStore) AtualizarStatus(id string, novoStatus string) (models.Leque, bool) {
	var lequeExistente models.Leque
	resultado := s.db.First(&lequeExistente, "id = ?", id)
	if resultado.Error != nil {
		return models.Leque{}, false
	}
	lequeExistente.Status = novoStatus
	err := s.db.Save(&lequeExistente).Error
	if err != nil {
		return models.Leque{}, false
	}
	return lequeExistente, true
}

func (s *LequeStore) Remover(id string) bool {
	var leque models.Leque
	resultado := s.db.First(&leque, "id = ?", id)
	if resultado.Error != nil {
		return false
	}
	err := s.db.Delete(&leque).Error
	if err != nil {
		return false
	}
	return true
}

func (s *LequeStore) Atualizar(id string, dadosNovos models.Leque) (models.Leque, error) {
	var lequeExistente models.Leque
	resultado := s.db.First(&lequeExistente, "id = ?", id)

	if resultado.Error != nil {
		return models.Leque{}, fmt.Errorf("leque não encontrado")
	}

	lequeExistente.Tipo = dadosNovos.Tipo
	lequeExistente.Numero = dadosNovos.Numero
	lequeExistente.Nome = dadosNovos.Nome
	lequeExistente.Status = dadosNovos.Status
	err := s.db.Save(&lequeExistente).Error
	if err != nil {
		return models.Leque{}, fmt.Errorf("erro ao atualizar leque: %w", err)
	}
	return lequeExistente, nil
}
