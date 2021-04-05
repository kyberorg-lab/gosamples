package main

import (
	"fmt"
	"log"
)

type Account struct {
	balance float64
}

func (a *Account) Balance() float64 {
	return a.balance
}

func (a *Account) Deposit(amount float64) {
	log.Printf("depositing: %f", amount)
	a.balance += amount
}

func (a *Account) Withdraw(amount float64) {
	if amount > a.balance {
		return
	}
	log.Printf("withdrawing: %f", amount)
	a.balance -= amount
}

func main() {
	acc := Account{}

	//start 10 goroutine
	for i := 0; i < 10; i++ {
		go func() {
			//each operates with Account
			for j := 0; j < 10; j++ {
				//sometimes withdraw "money"
				if j%2 == 1 {
					acc.Withdraw(50)
					continue
				}
				// sometimes deposit
				acc.Deposit(50)
			}
		}()
	}
	fmt.Scanln() //aka wait
	fmt.Println(acc.Balance())
}
