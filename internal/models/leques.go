package models

type Leque struct {
	ID     string `json:"id" gorm:"primaryKey"`
	AnelID string `json:"anel_id"`
	Tipo   string `json:"tipo"`
	Numero string `json:"numero"`
	Nome   string `json:"nome"`
	Status string `json:"status"`
}

func (Leque) TableName() string {
	return "leques"
}
