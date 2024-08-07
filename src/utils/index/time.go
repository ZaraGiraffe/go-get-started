package index_utils


import (
	"fmt"
)


func GetNumberOfIntervals(min_val int64, max_val int64, interval int64) int64 {
	return (max_val - min_val) / interval
}


var ErrOutOfIndexLeft = fmt.Errorf("index less than 0")
var ErrOutOfIndexRight = fmt.Errorf("index more than or equal to len(array)")

type TimestampUtil struct {
	min_timestamp int64
	length int
	interval int64
}


func InitTimestampUtil(min_timestamp int64, max_timestamp int64, interval int64) TimestampUtil {
	return TimestampUtil{
		min_timestamp,
		int((max_timestamp - min_timestamp) / interval),
		interval,
	}
}


func (t *TimestampUtil) TimestampToIndex(timestamp int64) (int, error) {
	if timestamp < t.min_timestamp {
		return -1, ErrOutOfIndexLeft
	}
	index := int((timestamp - t.min_timestamp) / t.interval)
	if index >= t.length {
		return -1, ErrOutOfIndexRight
	}
	return index, nil
}

func (t *TimestampUtil) IndexToTimestamp(index int) (int64, error) {
	if index >= t.length {
		return -1, ErrOutOfIndexLeft
	}
	return int64(index) * t.interval + t.min_timestamp, nil
}

func (t *TimestampUtil) GetLeftTimestamp(timestamp int64) int64 {
	/* get nearest left interval point */
	difference := timestamp - t.min_timestamp
	if difference % t.interval == 0 {
		return t.min_timestamp + difference 
	} else if difference > 0 {
		return t.min_timestamp + difference - (difference % t.interval)
	} else {
		return t.min_timestamp + difference - (difference % t.interval) - t.interval
	}
}

func (t *TimestampUtil) GetRightTimestamp(timestamp int64) int64 {
	/* get nearest right interval point */
	difference := timestamp - t.min_timestamp
	if difference % t.interval == 0 {
		return t.min_timestamp + difference
	} else if difference > 0 {
		return t.min_timestamp + difference - (difference % t.interval) + t.interval
	} else {
		return t.min_timestamp + difference - (difference % t.interval)
	}
}