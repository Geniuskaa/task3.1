package card

func TranslateMCC(Mcc string) (category string) {

	MCC := Mcc
	mcc := map[string]string{
		"5411": "Супермаркеты",
		"5812": "Рестораны",
	}

	Value, ok := mcc[MCC]
	switch ok {
	case true:
		category = Value
	case false:
		category = "Категория не указана"
	}
	return
}
