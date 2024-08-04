package main

import (
	"fmt"
	"slavka-test/src/constants"
)


func main() {
    path := constants.ConvertIntervalToPath(constants.DAY, constants.MEDIUM)
    fmt.Printf(path)
}