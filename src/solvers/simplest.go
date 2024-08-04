package solvers

import (
	"fmt"
	"slavka-test/src/data-utils"
	"slavka-test/src/file-utils"
	"slavka-test/src/index-utils"
	"slavka-test/src/structs"
	"slavka-test/src/constants"
)


func OneProcessSolver(interval_param constants.IntervalParam, size_param constants.SizeParam) {
	/* This is the solver that solves the task using one process without any optimizations */
	interval := int64(interval_param)

	user_data := file_utils.ReadUserData("big")
	market_data := file_utils.ReadMarketData("small")

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
	for ptr_user_data < len(user_data) && ptr_market_data < len(market_data) {
		
	}

}