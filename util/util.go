package util

import (
	"math/rand"
	"time"
)

type (
	// BankConnection interface definition
	BankConnection interface {
		GetBalanceByID(id int) (float64, error)
	}

	bankConnection struct {
		// some bank client internal connection details
	}
)

func (b bankConnection) GetBalanceByID(id int) (float64, error) {
	// This should be some side-effect operations like connections to DB and stuff
	// But for now we're going to return a random balance
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Float64() * 1000, nil
}

// GetBankConnection gets a connection to a bank
func GetBankConnection() BankConnection {
	return new(bankConnection)
}

// SumAccountBalances sum the amount of two accounts
func SumAccountBalances(accountA int, accountB int, client BankConnection) (float64, error) {
	balanceA, err := client.GetBalanceByID(accountA)
	if err != nil {
		return 0, err
	}

	balanceB, err := client.GetBalanceByID(accountB)
	if err != nil {
		return 0, err
	}

	return balanceA + balanceB, nil
}

// IsEven determines whether or not a number is even
func IsEven(number int) bool {
	if number%2 == 0 {
		return true
	}

	return false
}
