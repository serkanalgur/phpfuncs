package main

import (
	"fmt"

	"github.com/serkanalgur/phpfuncs"
)

func main() {
	fmt.Printf("Current Timestamp is: %d\n", phpfuncs.Time())

	phpfuncs.USleep(2000000)

	fmt.Printf("Current Timestamp After 2 seconds sleep is: %d\n", phpfuncs.Time())
}

// Expected result is based current timestamp and +2 seconds
// Give 2000000 for 2 seconds.
//  Current Timestamp is: 1607628607
//  Current Timestamp After 2 seconds sleep is: 1607628609
