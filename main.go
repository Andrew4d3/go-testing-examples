package main

import (
	"context"
	"fmt"

	"github.com/Andrew4d3/go-testing-examples/util"
)

func main() {
	fmt.Println("IsEven:", 4, util.IsEven(4))

	bankClient := util.GetBankConnection()
	balance, _ := bankClient.GetBalanceByID(123)

	ctx := context.WithValue(context.Background(), "foo", "bar")

	fmt.Println("Context Value in foo:", util.ExtractContextValue(ctx, "foo"))
	fmt.Println("Balance:", balance)
}
