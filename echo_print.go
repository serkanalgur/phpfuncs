package phpfuncs

import (
	"fmt"
)

// Echo - Output one or more strings
// Original : https://www.php.net/manual/en/function.echo.php
func Echo(vars ...interface{}){
	fmt.Print(vars...)
}

// Print - Output a string
// Original : https://www.php.net/manual/en/function.print.php
func Print(vars string) {
	fmt.Printf("%v\n",vars)
}

// PrintR - Prints human-readable information about a variable
// Original : https://www.php.net/manual/en/function.print-r.php
func PrintR(vars ...interface{}) {
	fmt.Print(vars ...)
}