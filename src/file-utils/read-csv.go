package file_utils

import (
	"encoding/csv"
	"os"
	"fmt"
	"strconv"
)

type UserData struct {
	user_id string
	currency string
	timestamp float64
	delta float64
}


type MarketData struct {
	symbol string
	timestamp float64
	price float64
}


func ReadUserData(data_type string) []UserData {
	file, _ := os.Open(fmt.Sprintf("./data/%s/user_data.csv", data_type))
	reader := csv.NewReader(file)
	all_rows, _ := reader.ReadAll()
	user_data := make([]UserData, len(all_rows)-1)
	for i, row := range all_rows {
		if i == 0 {
			continue
		}
		user_data[i-1].user_id = row[0]
		user_data[i-1].currency = row[1]
		user_data[i-1].timestamp, _ = strconv.ParseFloat(row[2], 64)
		user_data[i-1].delta, _ = strconv.ParseFloat(row[3], 64)
	} 
	return user_data
}


func ReadMarketData(data_type string) []MarketData {
	file, _ := os.Open(fmt.Sprintf("./data/%s/market_data.csv", data_type))
	reader := csv.NewReader(file)
	all_rows, _ := reader.ReadAll()
	market_data := make([]MarketData, len(all_rows)-1)
	for i, row := range all_rows {
		if i == 0 {
			continue
		}
		market_data[i-1].symbol = row[0]
		market_data[i-1].timestamp, _ = strconv.ParseFloat(row[1], 64)
		market_data[i-1].price, _ = strconv.ParseFloat(row[2], 64)
	}
	return market_data
}
