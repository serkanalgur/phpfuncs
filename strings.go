package phpfuncs

import (
	"log"
	"strconv"
	"strings"
)

// Addslashes - Quote string with slashes
// Original : https://www.php.net/manual/en/function.addslashes.php
// Returns a string with backslashes added before characters that need to be escaped.
func Addslashes(s string) string {

	l := len(s)
	r := make([]byte, l, l*2)
	for i := 0; i < l; i++ {

		switch s[i] {
		case '\'', '"', '\\':
			r = append(r, '\\')
		}

		r = append(r, s[i])
	}

	return string(r)
}

// Bin2hex - Convert binary data into hexadecimal representation
// Original : https://www.php.net/manual/en/function.bin2hex.php
// Returns an ASCII string containing the hexadecimal representation of str. The conversion is done byte-wise with the high-nibble first.
func Bin2hex(s string) string {
	bin, err := strconv.ParseInt(s,2,64)
	if err != nil {
		log.Fatal(err)
	}
	return strconv.FormatInt(bin,16)
}

// Explode - Split a string by a string
// Original : https://www.php.net/manual/en/function.explode.php
// Returns an array of strings, each of which is a substring of string formed by splitting it on boundaries formed by the string delimiter.
func Explode(s,set string) []string {
	return strings.Split(s,set)
}

// Join - Alias o Implode
// Original : https://www.php.net/manual/en/function.join.php
// Alias o Implode
func Join(set string, s []string) string {
	return Implode(set,s)
}

// Ltrim - Strip whitespace (or other characters) from the beginning of a string
// Original : https://www.php.net/manual/en/function.ltrim.php
// Strip whitespace (or other characters) from the beginning of a string.
func Ltrim(s, set string) string{
	if set == "" { set = " "}
	return strings.TrimLeft(s, set)
}

// Rtrim - Strip whitespace (or other characters) from the end of a string
// Original : https://www.php.net/manual/en/function.rtrim.php
// This function returns a string with whitespace (or other characters) stripped from the end of str.
func Rtrim(s, set string) string {
	if set == "" { set = " "}
	return strings.TrimRight(s, set)
}

// Strtolower - Make a string lowercase
// Original : https://www.php.net/manual/en/function.strtolower.php
// Returns string with all alphabetic characters converted to lowercase.
func Strtolower(s string) string {
	return strings.ToLower(s)
}


// Strtoupper - Make a string uppercase
// Original : https://www.php.net/manual/en/function.strtoupper.php
// Returns string with all alphabetic characters converted to uppercase.
func Strtoupper(s string) string {
	return strings.ToUpper(s)
}

// MbStrtolower - Make a string lowercase
// Original : https://www.php.net/manual/en/function.mb-strtolower.php
// Returns str with all alphabetic characters converted to lowercase.
func MbStrtolower(s string) string {
	return Strtolower(s)
}

// MbStrtoupper - Make a string uppercase
// Original : https://www.php.net/manual/en/function.mb-strtoupper.php
// Returns str with all alphabetic characters converted to uppercase.
func MbStrtoupper(s string) string {
	return Strtoupper(s)
}

// Trim - Strip whitespace (or other characters) from the beginning and end of a string
// Original : https://www.php.net/manual/en/function.trim.php
// This function returns a string with whitespace stripped from the beginning and end of str.
func Trim(s, set string) string {
	if set == "" {return strings.TrimSpace(s)}
	return strings.Trim(s, set)
}
