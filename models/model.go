package models

func GetModels() []interface{} {
	return []interface{}{
		&User{},
		&RebateLog{},
		&GoodShareProfit{},
	}
}
