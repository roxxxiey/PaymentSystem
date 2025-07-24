package handlers

import (
	"PaymentSystem/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

func GetBalance(address string) (float32, error) {
	if address == "" {
		return 0, errors.New("address is empty")
	}

	addressUUID, err := uuid.Parse(address)
	if err != nil {
		return 0, fmt.Errorf("invalid address UUID: %v", err)
	}

	var wallet models.Wallet
	if err := DB.Where("address = ?", addressUUID).First(&wallet).Error; err != nil {
		return 0, fmt.Errorf("failed to find wallet: %v", err)
	}

	return wallet.Amount, nil
}
