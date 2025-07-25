package handlers

import (
	"PaymentSystem/models"
	"errors"
	"strconv"
)

func GetLastTransactions(count string) ([]*models.Transaction, error) {
	countInt, err := strconv.Atoi(count)
	if err != nil {
		return nil, errors.New("parse string to int: " + err.Error())
	}

	if countInt < 1 {
		return nil, errors.New("count transactions sube zero")
	}

	var transactions []*models.Transaction

	result := DB.
		Order("time_trans DESC").
		Limit(countInt).
		Find(&transactions)

	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}
