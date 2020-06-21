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

func SumByMCC(transaction []Transaction, mcc []string) (sum int64) {
	length := len(transaction)
	if (contains(mcc, transaction[length - 1].MCC) == true) {
		sum = transaction[length - 1].SumOfTransaction
	}
	return
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

