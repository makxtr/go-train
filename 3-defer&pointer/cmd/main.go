package main

import "fmt"

type Account struct {
	Balance int
}

func main() {
	initialBalance := 1000
	account := &Account{Balance: initialBalance}

	defer printBalance("Изначальный баланс", account.Balance)
	defer printBalance("Текущий баланс", account.Balance)
	defer printAccountBalance("Указатель на баланс", account)

	account.Balance += 500
	updateBalance(account, 200)
	account = &Account{Balance: 300}
}

func updateBalance(account *Account, amount int) {
	account.Balance -= amount
}

func printBalance(label string, balance int) {
	fmt.Printf("%s: %d\n", label, balance)
}

func printAccountBalance(label string, account *Account) {
	fmt.Printf("%s: %d\n", label, account.Balance)
}
