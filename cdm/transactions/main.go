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
		Id:               72,
		SumOfTransaction: 1_203_91,
		Date:             1592657460,
		MCC:              "5852",
		Status:           "Обработано",
	}
	transaction2 := card.Transaction{
		Id:               83,
		SumOfTransaction: 1_273_71,
		Date:             1592657560,
		MCC:              "5411",
		Status:           "Обработано",
	}

	transaction3 := card.Transaction{
		Id:               94,
		SumOfTransaction: 1_563_91,
		Date:             1592743860,
		MCC:              "5812",
		Status:           "Обработано",
	}

	card.AddTransaction(master, transaction1)
	card.SumByMCC(transaction1, transaction1.MCC)

	fmt.Println(card.SumByMCC(transaction1, transaction1.MCC))
	fmt.Println(master)

	category := card.TranslateMCC(master.Transactions[0].MCC)
	fmt.Println(category)

	card.AddTransaction(master, transaction2)
	card.AddTransaction(master, transaction3)


	fmt.Println(card.LastNTransactions(master, 3))

}
