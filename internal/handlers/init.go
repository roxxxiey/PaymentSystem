package handlers

import (
	"PaymentSystem/models"
	"fmt"
	"github.com/google/uuid"
)

func AutoMigrateAndInit() error {
	if DB == nil {
		return fmt.Errorf("database is not set")
	}

	err := DB.AutoMigrate(&models.Wallet{}, &models.Transaction{})
	if err != nil {
		return fmt.Errorf("auto migration failed: %v", err)
	}

	// Проверим, есть ли кошельки уже
	var count int64
	if err := DB.Model(&models.Wallet{}).Count(&count).Error; err != nil {
		return fmt.Errorf("count wallets failed: %v", err)
	}

	if count == 0 {
		// Создаем 10 кошельков
		for i := 0; i < 10; i++ {
			w := models.Wallet{
				Address: uuid.New().String(),
				Amount:  100.0,
			}
			if err := DB.Create(&w).Error; err != nil {
				return fmt.Errorf("failed to create wallet: %v", err)
			}
		}
		fmt.Println("10 wallets with 100 balance created")
	} else {
		fmt.Println("Wallets already exist, skipping init")
	}

	return nil
}
