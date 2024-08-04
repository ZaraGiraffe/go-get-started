package file_utils

import (
	"os"
	"fmt"
	"slavka-test/src/structs"
	"slavka-test/src/index-utils"
)



func WriteAnswerData(answer structs.AnswerData, users_map map[string]int, time_util index_utils.TimestampUtil) {
	user_ids := index_utils.GetInverseStringMap(users_map)

	file, _ := os.Create("./")
	file.WriteString("user_id,minimum_balance,maximum_balance,average_balance,start_timestamp\n")

	for i, data_unit_1d := range answer.Content {
		for j, data_unit := range data_unit_1d {
			timestamp, _ := time_util.IndexToTimestamp(j)
			output_string := fmt.Sprintf("%s,%f,%f,%f,%d",
			 user_ids[i], data_unit.MinimumBalance, data_unit.MaximumBalance, data_unit.AverageBalance, timestamp)
			file.WriteString(output_string)
		}
	}
}
