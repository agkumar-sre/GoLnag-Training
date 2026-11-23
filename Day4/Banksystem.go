package main

import "fmt"

type Account struct {
	Name string
	Balance float64
}

func (a *Account) Deposit(amount float64){
	if amount <= 0 {
		fmt.Println("Invalid deposit amount")
		return
	}

	a.Balance += amount
	fmt.Println("Deposite Sucessful")
}

func(a *Account) Withdraw(amount float64){
	if amount <= 0 {
		fmt.Println("Invalid Withdraw amount")
		return
	}

	if amount > a.Balance {
		fmt.Println("Insufficient balance")
		return
	}

	a.Balance -= amount
	fmt.Println("Withdrawl Successful")
}

func main(){
	account := Account{
		Name: "Amit",
		Balance: 1000,
	}

	fmt.Println("Account Holder:", account.Name)
	fmt.Println("Initial Balance:", account.Balance)

	account.Deposit(4000)
	fmt.Println("Balance after deposit:", account.Balance)

	account.Withdraw(3000)
	fmt.Println("Balance after withdrawal:", account.Balance)

	

}
