package card

type Transaction struct {
	Id               int64
	SumOfTransaction int64
	Date             int64 // в формате Unix Timeship
	MCC              string
	Status           string
}

type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Icon         string
	Transactions []Transaction
}

func AddTransaction(card *Card, transaction Transaction) *Card {
	card.Transactions = append(card.Transactions, transaction)
	return card
}

func SumByMCC(transaction Transaction, mcc string) (sum int64) {
	switch mcc {
	case "5411":
		sum = transaction.SumOfTransaction
	case "5812":
		sum = transaction.SumOfTransaction
	case "", " ":
		sum = 0
	default:
		sum = 0
	}
	return sum
}

func LastNTransactions(card *Card, n int) []Transaction {
	length := len(card.Transactions)
	i := 0
	//var m int
	var copyOfTransactions []Transaction
	if (length < n || length == n) {
		copyOfTransactions = make([]Transaction, length)
		for (length > 0) {
			copyOfTransactions[i] = card.Transactions[length - 1]
			i++
			length--
		}
		//m = copy(copyOfTransactions, card.Transactions)
	} else {
		copyOfTransactions = make([]Transaction, n)
		for (n > 0) {
			copyOfTransactions[i] = card.Transactions[length - 1]
			i++
			length--
			n--
		}
		//m = copy(copyOfTransactions, card.Transactions[(len(card.Transactions) - n):])
	}
	//fmt.Println(m)
	return copyOfTransactions
}