package main

import (
	"fmt"
	card "github.com/Geniuskaa/task3.1/pkg/transactions"
)

func main() {
	historyOfTransactions := []card.Transaction{}
	master := &card.Card{
		Id:           2389,
		Issuer:       "Visa",
		Balance:      23_000_00,
		Currency:     "RUB",
		Number:       "2482",
		Icon:         "https://.....",
		Transactions: historyOfTransactions,
	}

	fmt.Println(master.Transactions)
	fmt.Println(master)

	transaction1 := card.Transaction{
		Id:               82,
		SumOfTransaction: 1_203_91,
		Date:             1592657460,
		MCC:              "5812",
		Status:           "Обработано",
	}

	card.AddTransaction(master, transaction1)
	card.SumByMCC(transaction1, transaction1.MCC)

	fmt.Println(card.SumByMCC(transaction1, transaction1.MCC))
	fmt.Println(master)

	category := card.TranslateMCC(master.Transactions[0].MCC)
	fmt.Println(category)


}
