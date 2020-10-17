package main

import (
	"fmt"

	"github.com/Andrew4d3/go-testing-examples/util"
)

func main() {
	fmt.Println("IsEven:", 4, util.IsEven(4))

	bankClient := util.GetBankClient()
	balance, _ := bankClient.GetBalanceByID(123)

	fmt.Println("Balance:", balance)
}
