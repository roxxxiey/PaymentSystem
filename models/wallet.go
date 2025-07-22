package models

type Wallet struct {
	IP      int     `gorm:"primaryKey" json:"id"`
	Address string  `json:"address"`
	Amount  float32 `json:"amount"`
}
