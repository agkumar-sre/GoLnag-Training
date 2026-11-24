package main 

import(
	"fmt"
	"banking-app/bank"
)

func main(){
	account := bank.Account{Balance: 1000}

	account.Deposit(5000)
	fmt.Println("Balance:", account.Balance)

	err := account.Withdrawn(2000)

	if err != nil{
		fmt.Println("Error:", err)
	}

	fmt.Println("Final Balance:", account.Balance)
}
