package main

import (
	"fmt"

	"github.com/serkanalgur/phpfuncs"
)

func main(){
	Encbase64 := phpfuncs.Base64Encode("Hello World")
	fmt.Printf("Base64 encoded \"Hello World\" is: %v\n",Encbase64)
}

// Expected : Base64 encoded "Hello World" is: SGVsbG8gV29ybGQ=