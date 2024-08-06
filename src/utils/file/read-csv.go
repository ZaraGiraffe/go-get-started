package file_utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"slavka-test/src/constants"
	"slavka-test/src/structs"
	"strconv"
)




func ReadUserData(size_param constants.SizeParam) []structs.UserData {
	/* This function reads the user data to the special type */
	file, _ := os.Open(fmt.Sprintf("./data/%s/user_data.csv", string(size_param)))
	reader := csv.NewReader(file)
	all_rows, _ := reader.ReadAll()
	user_data := make([]structs.UserData, len(all_rows)-1)
	for i, row := range all_rows {
		if i == 0 {
			continue
		}
		user_data[i-1].UserId = row[0]
		user_data[i-1].Currency = row[1]
		user_data[i-1].Timestamp, _ = strconv.ParseInt(row[2], 10, 64)
		user_data[i-1].Delta, _ = strconv.ParseFloat(row[3], 64)
	} 
	return user_data
}


func ReadMarketData(size_param constants.SizeParam) []structs.MarketData {
	/* This function reads the market data to the special type */
	file, _ := os.Open(fmt.Sprintf("./data/%s/market_data.csv", string(size_param)))
	reader := csv.NewReader(file)
	all_rows, _ := reader.ReadAll()
	market_data := make([]structs.MarketData, len(all_rows)-1)
	for i, row := range all_rows {
		if i == 0 {
			continue
		}
		market_data[i-1].Symbol = row[0]
		market_data[i-1].Timestamp, _ = strconv.ParseInt(row[1], 10, 64)
		market_data[i-1].Price, _ = strconv.ParseFloat(row[2], 64)
	}
	return market_data
}
