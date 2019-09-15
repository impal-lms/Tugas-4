package main

import "fmt"

type StatusCode int

// StatusCode enum
const (
	Valid StatusCode = iota
	Invalid
)

type Account struct {
	Status       StatusCode
	AccNumber    int
	AmountOfSale int
}

var Accounts = make(map[int]Account)

// dummy data
func initAccount() {
	for i := 0; i < 100; i++ {
		var status StatusCode
		if i%2 == 0 {
			status = Valid
		} else {
			status = Invalid
		}
		Accounts[i] = Account{
			status,
			i,
			i * 100000,
		}
	}
}

func main() {
	initAccount()
	fmt.Println("your account number: ")
	var accNumber int
	fmt.Scanln(&accNumber)

	account, ok := Accounts[accNumber]
	if !ok {
		fmt.Println("ERR: account not found")
	} else if account.Status != Valid {
		fmt.Println("ERR: account is not valid")
	} else {
		fmt.Printf("amount of sale: %d\n", account.AmountOfSale)
	}

}
