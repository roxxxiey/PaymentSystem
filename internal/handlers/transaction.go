package handlers

import (
	"PaymentSystem/models"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func SetDB(db *gorm.DB) {
	DB = db
}

func SendMoney(fromTrans, toTrans string, amount float32) error {

	if fromTrans == toTrans {
		return errors.New("can't send money to same address")
	}
	if fromTrans == "" || toTrans == "" {
		return errors.New("can't send money to nil address")
	}
	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	fromUUID, err := uuid.Parse(fromTrans)
	if err != nil {
		return fmt.Errorf("invalid sender UUID: %w", err)
	}
	toUUID, err := uuid.Parse(toTrans)
	if err != nil {
		return fmt.Errorf("invalid receiver UUID: %w", err)
	}

	var fromWallet, toWallet models.Wallet

	tx := DB.Begin()

	if err = tx.Where("address = ?", fromTrans).First(&fromWallet).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("sender wallet not found: %w", err)
	}

	if fromWallet.Amount < amount {
		tx.Rollback()
		return errors.New("insufficient funds")
	}

	if err := tx.Where("address = ?", toTrans).First(&toWallet).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("receiver wallet not found: %w", err)
	}

	fromWallet.Amount -= amount
	if err := tx.Save(&fromWallet).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update sender wallet: %w", err)
	}

	toWallet.Amount += amount
	if err := tx.Save(&toWallet).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update receiver wallet: %w", err)
	}

	transaction := models.Transaction{
		FromTrans: fromUUID,
		ToTrans:   toUUID,
		Amount:    amount,
		TimeTrans: time.Now(),
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create transaction: %w", err)
	}

	return tx.Commit().Error
}
