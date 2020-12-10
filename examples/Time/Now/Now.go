package main

import (
	"fmt"

	"github.com/serkanalgur/phpfuncs"
)

func main(){
	currentTime := phpfuncs.Now()
	fmt.Printf("Current Timestamp is: %d",currentTime)
}

// Expected : Current Timestamp is : "current timestamp"