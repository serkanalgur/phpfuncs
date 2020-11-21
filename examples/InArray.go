package main

import (
	"fmt"
	"strings"

	"github.com/serkanalgur/phpfuncs"
)

  func main(){

    test := strings.Split("Denee için serkan algur kişisinden birşeyler seçmeliyiz."," ")

    arrayo := phpfuncs.InArray("serkaaaan",test)

    fmt.Printf("Result is : %v\n",arrayo)
  }