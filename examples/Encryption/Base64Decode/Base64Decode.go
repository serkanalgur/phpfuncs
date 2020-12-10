package main

import (
	"fmt"

	"github.com/serkanalgur/phpfuncs"
)

func main() {
	Decbase64, _ := phpfuncs.Base64Decode("SGVsbG8gV29ybGQ=")
	fmt.Printf("Base64 decoded \"SGVsbG8gV29ybGQ=\" is: %v\n", Decbase64)
}

// Expected : Base64 encoded "SGVsbG8gV29ybGQ=" is: Hello World
