package data_utils


import (
	"slavka-test/src/structs"
	"golang.org/x/exp/maps"
)


func GetMinMaxTimestamp(user_data []structs.UserData, market_data []structs.MarketData, interval int64) (int64, int64) {
	min_val_user_data := user_data[0].Timestamp / interval * interval
	max_val_user_data := (user_data[len(user_data)-1].Timestamp + interval - 1) / interval * interval
	min_val_market_data := market_data[0].Timestamp / interval * interval
	max_val_market_data := (market_data[len(market_data)-1].Timestamp + interval - 1) / interval * interval
	return max(min_val_user_data, min_val_market_data), max(max_val_user_data, max_val_market_data)
}


func GetListOfUsers(user_data []structs.UserData) []string {
	user_map := make(map[string]bool)
	for _, row := range user_data {
		user_map[row.UserId] = true
	}
	return maps.Keys(user_map)
}

func GetListOfCurrenciesAndPrices(user_data []structs.UserData) ([]string, []string) {
	/* Returns the list of currencies and currencies + USD */
	user_map := make(map[string]bool)
	for _, row := range user_data {
		user_map[row.Currency] = true
	}
	currencies := maps.Keys(user_map)
	prices := make([]string, len(currencies))
	for i := 0; i < len(currencies); i++ {
		prices[i] = prices[i] + "USD"
	}
	return currencies, prices
}
