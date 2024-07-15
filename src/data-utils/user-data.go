package data_utils


import (
	"slavka-test/src/file-utils"
)


func GetMinMaxTimestamp(user_data []file_utils.UserData, market_data []file_utils.MarketData, interval int64) (int64, int64) {
	min_val_user_data := user_data[0].Timestamp / interval * interval
	max_val_user_data := (user_data[len(user_data)-1].Timestamp + interval - 1) / interval * interval
	min_val_market_data := market_data[0].Timestamp / interval * interval
	max_val_market_data := (market_data[len(market_data)-1].Timestamp + interval - 1) / interval * interval
	return max(min_val_user_data, min_val_market_data), max(max_val_user_data, max_val_market_data)
}