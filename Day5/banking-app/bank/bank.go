package bank

import "errors"

type Account struct {
	Balance float64
}

func (a *Account) Deposit(amount float64){
	a.Balance += amount
}

func (a *Account) Withdrawn(amount float64) error{
	if amount > a.Balance {
		return errors.New("insufficient balance")
	}

	a.Balance -= amount
	return nil
}

