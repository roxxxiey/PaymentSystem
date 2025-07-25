package models

import (
	"github.com/google/uuid"
	"time"
)

type Transaction struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	FromTrans uuid.UUID `json:"from_trans"`
	ToTrans   uuid.UUID `json:"to_trans"`
	Amount    float32   `json:"amount"`
	TimeTrans time.Time `json:"time_trans"`
}
