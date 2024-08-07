package solvers

import (
	"slavka-test/src/utils/data"
	"slavka-test/src/utils/file"
	"slavka-test/src/utils/index"
	"slavka-test/src/structs"
	"slavka-test/src/constants"
)


func OneProcessSolver(interval_param constants.IntervalParam, size_param constants.SizeParam) {
	/* This is the solver that solves the task using one process without any optimizations */
	interval := int64(interval_param)

	user_data := file_utils.ReadUserData(size_param)
	market_data := file_utils.ReadMarketData(size_param)

	min_timestamp, max_timestamp := data_utils.GetMinMaxTimestamp(user_data, market_data, interval)
	users := data_utils.GetListOfUsers(user_data)
	currencies, prices := data_utils.GetListOfCurrenciesAndPrices(user_data)
	answer := structs.InitAnswerData(len(users), int((max_timestamp - min_timestamp) / interval), interval_param)
	cur_users_data := structs.InitUserCurrentData(len(users), len(currencies))

	currency_map := index_utils.GetStringMap(currencies)
	prices_map := index_utils.GetStringMap(prices)
	users_map := index_utils.GetStringMap(users)

	cur_currencies_prices := make([]float64, len(prices))
	cur_currencies_prices[prices_map["USDUSD"]] = 1.0

	time_util := index_utils.InitTimestampUtil(min_timestamp, max_timestamp, interval)

	ptr_user_data := 0
	ptr_market_data := 0
	next_timestamp := min_timestamp

	for ptr_user_data < len(user_data) || ptr_market_data < len(market_data) {

		user_data_time := constants.INF64
		if ptr_user_data < len(user_data) {
			user_data_time = user_data[ptr_user_data].Timestamp
		}
		market_data_time := constants.INF64
		if ptr_market_data < len(market_data) {
			market_data_time = market_data[ptr_market_data].Timestamp
		}
		
		cur_time := min(user_data_time, market_data_time)
		cur_timestamp := time_util.GetLeftTimestamp(cur_time)
		for ; cur_timestamp >= next_timestamp; next_timestamp += interval {
			index_timestamp, _ := time_util.TimestampToIndex(next_timestamp)
			for _, v := range users_map {
				answer_unit := &answer.Content[v][index_timestamp]
				answer_unit.AverageBalance = cur_users_data[v].BalanceInUSD
				answer_unit.MaximumBalance = cur_users_data[v].BalanceInUSD
				answer_unit.MinimumBalance = cur_users_data[v].BalanceInUSD
			}
		}

		updateUser := func(user_id int) {
			cur_user_data := &cur_users_data[user_id]
			new_balance := cur_user_data.RecalculateUSDBalance(cur_currencies_prices)
			index_timestamp, err := time_util.TimestampToIndex(cur_timestamp)
			if err == nil {
				answer_unit := &answer.Content[user_id][index_timestamp]
				if new_balance > answer_unit.MaximumBalance {
					answer_unit.MaximumBalance = new_balance
				}
				if new_balance < answer_unit.MinimumBalance {
					answer_unit.MinimumBalance = new_balance
				}
				answer_unit.AverageBalance += (new_balance - cur_user_data.BalanceInUSD) * 
					float64(time_util.GetRightTimestamp(cur_time) - cur_time) / float64(interval)
			}
			cur_user_data.BalanceInUSD = new_balance
		}

		if user_data_time < market_data_time {
			user_data_unit := &user_data[ptr_user_data]
			user_id := users_map[user_data_unit.UserId]
			cur_user_data := cur_users_data[user_id]
			currency_id := currency_map[user_data_unit.Currency]
			cur_user_data.BalanceInCurrency[currency_id] += user_data_unit.Delta
			updateUser(user_id)
			ptr_user_data += 1
		} else {
			market_data_unit := &market_data[ptr_market_data]
			price_id := prices_map[market_data_unit.Symbol]
			cur_currencies_prices[price_id] = market_data_unit.Price
			for user_id := 0; user_id < len(cur_users_data); user_id++ {
				updateUser(user_id)
			}
			ptr_market_data += 1
		}
	}

	file_path := constants.ConvertParamsToPath(interval_param, size_param)
	file_utils.WriteAnswerData(answer, users_map, time_util, file_path)
}