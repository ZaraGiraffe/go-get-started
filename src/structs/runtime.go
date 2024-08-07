package structs

import (
	"slavka-test/src/constants"
)


type UserCurrentData struct {
	BalanceInCurrency []float64
	BalanceInUSD float64
}

type AnswerDataUnit struct {
	MinimumBalance float64
	MaximumBalance float64
	AverageBalance float64
}

type AnswerData = struct {
	/* For every user, For every timestamp saves the answer */
	Content [][]AnswerDataUnit
	Type constants.IntervalParam
}


func InitAnswerData(num_users int, num_timestamp int, interval constants.IntervalParam) AnswerData {
	content := make([][]AnswerDataUnit, num_users)
	for i := 0; i < num_users; i++ {
		content[i] = make([]AnswerDataUnit, num_timestamp)
	}
	answer_data := AnswerData {
		content, 
		interval,
	}
	return answer_data
}


func InitUserCurrentData(num_users int, num_currencies int) []UserCurrentData {
	cur_data := make([]UserCurrentData, num_users)
	for i := 0; i < num_users; i++ {
		cur_data[i] = UserCurrentData {
			make([]float64, num_currencies),
			0.0,
		}
	}
	return cur_data
}



func (u *UserCurrentData) RecalculateUSDBalance(current_currency_prices []float64) float64 {
	var new_balance float64 = 0.0
	for i := 0; i < len(u.BalanceInCurrency); i++ {
		new_balance += u.BalanceInCurrency[i] * current_currency_prices[i]
	}
	return new_balance
}