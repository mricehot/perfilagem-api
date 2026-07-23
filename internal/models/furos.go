package models

type Furo struct {
	ID               string  `json:"id" gorm:"primaryKey"`
	LequeID          string  `json:"leque_id"`
	Numero           string  `json:"numero"`
	MetragemEsperada float64 `json:"metragem_esperada"`
	MetragemReal     float64 `json:"metragem_real"`
	Situacao         string  `json:"situacao"`
}

func (Furo) TableName() string {
	return "furos"
}
