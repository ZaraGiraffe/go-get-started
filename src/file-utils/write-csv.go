package file_utils

import (
	"slavka-test/src/structs"
	"slavka-test/src/index_utils"
	"strconv"
	"encodings/csv"
)




func convert_answer_to_strings(answer structs.AnswerData, users_map map[string]int, time_util index_utils.TimestampUtil) [][]string {
	/* Converts the special AnswerData object to the object list of list of strings */
	users_list := index_utils.GetInverseStringMap(users_map)
	dim1 := len(answer)
	dim2 := 5
	strings := make([][]string, dim1+1)
	for i := range dim1+1 {
		strings[i] = make([]string, dim2)
	}
	strings[0] = []string{"user_id","minimum_balance","maximum_balance","average_balance","start_timestamp"}
	for i := range dim1 {
		user_id := users_list[i]
		for j := range dim2 {
			timestamp := time_util.IndexToTimestamp(j)
			timestamp_str := strconv.FormatInt(timestamp, 10)
			minimum_str := strconv.FormatFloat(answer[i][j].MinimumBalance, 'f', -1, 64)
			maximum_str := strconv.FormatFloat(answer[i][j].MaximumBalance, 'f', -1, 64)
			average_str := strconv.FormatFloat(answer[i][j].AverageBalance, 'f', -1, 64)
			strings[i+1] = []string{user_id, minimum_str, maximum_str, average_str, timestamp_str}
		}
	}
	return strings
}


func WriteAnswerData(answer structs.AnswerData, users_map map[string]int, time_util index_utils.TimestampUtil) {
	/* Write the AnswerData structure to the file */
	strings := convert_answer_to_strings(answer, users_map, time_util)
	writer := csv.NewWriter()

}
