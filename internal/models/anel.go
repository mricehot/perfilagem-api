package models

type Anel struct {
	ID    string `json:"id" gorm:"primaryKey"`
	Nome  string `json:"nome"`
	Ativo bool   `json:"ativo"`
}

func (Anel) TableName() string {
	return "aneis"
}
