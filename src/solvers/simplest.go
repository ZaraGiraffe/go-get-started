package solvers

import (
	"fmt"
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
	answer := structs.InitAnswerData(len(users), int((min_timestamp - max_timestamp) / interval), interval_param)
	cur_user_data := structs.InitUserCurrentData(len(users))

	currency_map := index_utils.GetStringMap(currencies)
	prices_map := index_utils.GetStringMap(prices)
	users_map := index_utils.GetStringMap(users)

	time_util := index_utils.InitTimestampUtil(min_timestamp, max_timestamp, interval)

	ptr_user_data := 0
	ptr_market_data := 0
	next_timestamp := min_timestamp

	for ptr_user_data < len(user_data) && ptr_market_data < len(market_data) {

		user_data_time := constants.INF64
		if ptr_user_data < len(user_data) {
			user_data_time = user_data[ptr_user_data].Timestamp
		}
		market_data_time := constants.INF64
		if ptr_market_data < len(market_data) {
			market_data_time = market_data[ptr_market_data].Timestamp
		}

		cur_timestamp := time_util.GetLeftTimestamp(min(user_data_time, market_data_time))
		for ; cur_timestamp >= next_timestamp; next_timestamp += interval {
			index_timestamp, _ := time_util.TimestampToIndex(cur_timestamp)
			for _, v := range users_map {
				answer.Content[v][index_timestamp].AverageBalance = cur_user_data.BalanceInUSD
				answer.Content[v][index_timestamp].MaximumBalance = cur_user_data.BalanceInUSD
				answer.Content[v][index_timestamp].MinimumBalance = cur_user_data.BalanceInUSD
			}
		}
		
		
	}
}