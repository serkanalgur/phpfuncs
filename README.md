# phpfuncs

Php Functions in Go

## Installation

```bash
  go get github.com/serkanalgur/phpfuncs
```

Description and other info will be here.

## Functions

- [InArray (PHP in_array)](#inarray-php-in_array)



### InArray (PHP in_array)

#### Usage

  ```go

    //needle str|int
    //haystack array []string, []int

    phpfuncs.InArray(needle,haystack)
    //Return will be true or false like in PHP
  ```

#### Example

You can see in [examples folder](examples/InArray.go)

```go
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
````
