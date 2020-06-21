package main

import (
	"fmt"
	"github.com/Geniuskaa/task3.1/pkg/card"
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

	mcc := []string {"5411", "5812"}

	card.AddTransaction(master, transaction1)
	card.SumByMCC(master.Transactions, mcc)

	fmt.Println(card.SumByMCC(master.Transactions, mcc))
	fmt.Println(master)

}
