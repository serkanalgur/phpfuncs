package main

import (
	"fmt"

	"github.com/serkanalgur/phpfuncs"
)

func main(){
	currentTime := phpfuncs.Time()
	fmt.Printf("Current Timestamp is: %d",currentTime)
}

// https://play.golang.org/p/xzhpqD0nNJh