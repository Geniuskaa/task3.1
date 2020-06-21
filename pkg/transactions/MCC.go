package card

const errCategoryNotFound = "Категория не указана"
func TranslateMCC(Mcc string) string {
	MCC := Mcc
	mcc := map[string]string{
		"5411": "Супермаркеты",
		"5812": "Рестораны",
	}
	if value, ok := mcc[MCC]; ok {
		return value
	}
	return errCategoryNotFound;
}
