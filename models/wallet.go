package models

type Wallet struct {
	IP      int     `gorm:"primaryKey" json:"id"`
	Address string  `json:"address"gorm:"uniqueIndex""`
	Amount  float32 `json:"amount" gorm:"not null;default:0"`
}
