package models

type Wallet struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Firstname  string `json:"firstname"`
	Lastname string `json:"lastname"`
	Amount float32 `json:"amount"`
}
