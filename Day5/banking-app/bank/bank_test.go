package bank

import "testing"

func TestOverdraft(t *testing.T){
	account := Account{Balance: 1000}

	err := account.Withdrawn(1500)

	if err == nil {
		t.Fatal("Expected overdraft error")
	}

	if account.Balance != 1000 {
		t.Errorf("Balance changed unexpected: %v", account.Balance)
	}
}