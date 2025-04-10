package main

import (
	"errors"
)

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

func updateBalance(customer *customer, transaction transaction) error {
	switch transaction.transactionType {
	case transactionDeposit:
		customer.balance += transaction.amount
		return nil

	case transactionWithdrawal:
		if customer.balance < transaction.amount {
			return errors.New("insufficient funds")
		}

		customer.balance -= transaction.amount
		return nil

	default:
		return errors.New("unknown transaction type")
	}
}
