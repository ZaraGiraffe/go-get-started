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


func ReadUserData(data_type string) []UserData {
	file, err := os.Open(fmt.Sprintf("./data/%s/user_data.csv", data_type))
	if err != nil {
		fmt.Println("error in opening the file")
	}
	reader := csv.NewReader(file)
	all_rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("error in the reader")
	}
	user_data := make([]UserData, len(all_rows)-1)
	for i, row := range all_rows {
		if i == 0 {
			continue
		}
		user_data[i-1].user_id = row[0]
		user_data[i-1].currency = row[1]
		user_data[i-1].timestamp, err = strconv.ParseFloat(row[2], 64)
		if err != nil {
			fmt.Printf("error in the timestamp, line %d, %e", i, err)
		}
		user_data[i-1].delta, err = strconv.ParseFloat(row[3], 64)
		if err != nil {
			fmt.Printf("error in the delta, line %d, %e", i, err)
		}
	} 
	return user_data
}