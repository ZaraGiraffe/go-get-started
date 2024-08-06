package constants

import (
	"fmt"
)


type IntervalParam int64

const (
	HOUR IntervalParam = 60 * 60
	DAY IntervalParam = HOUR * 24
	MONTH IntervalParam = DAY * 30
)

type SizeParam string

const (
	BIG SizeParam = "big"
	SMALL SizeParam = "small"
	MEDIUM SizeParam = "medium"
)


func ConvertIntervalToPath(interval IntervalParam, size SizeParam) string {
	var interval_string string
	if interval == HOUR {
		interval_string = "1h"
	} else if interval == DAY {
		interval_string = "1d"
	} else {
		interval_string = "30d"
	}

	path := fmt.Sprintf("./data/%s/example-output/bars-%s.csv", size, interval_string)
	return path
}


const INF64 int64 = 9223372036854775807