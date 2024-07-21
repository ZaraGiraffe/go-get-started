package structs

type UserCurrentData struct {
	BalanceInCurrency []float64
	BalanceInUSD float64
}

type AnswerDataUnit struct {
	MinimumBalance float64
	MaximumBalance float64
	AverageBalance float64
}

type AnswerData = [][]AnswerDataUnit

func InitAnswerData(num_users int, num_timestamp int) AnswerData {
	answer_data := make([][]AnswerDataUnit, num_users)
	for i := range num_users {
		answer_data[i] = make([]AnswerDataUnit, num_timestamp)
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

