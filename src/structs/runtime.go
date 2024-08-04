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
	for i := range num_users {
		content[i] = make([]AnswerDataUnit, num_timestamp)
	}
	answer_data := AnswerData {
		content, 
		interval,
	}
	return answer_data
}


func InitUserCurrentData(num_users int) UserCurrentData {
	cur_data := UserCurrentData{
		BalanceInCurrency: make([]float64, num_users),
		BalanceInUSD: 0,
	}
	return cur_data
}
