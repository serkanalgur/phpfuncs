package phpfuncs

import (
	"html"
	"log"
	"strconv"
	"strings"
)

// Addslashes - Quote string with slashes
//
// Original : https://www.php.net/manual/en/function.addslashes.php
//
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

// Bin2Hex - Convert binary data into hexadecimal representation
//
// Original : https://www.php.net/manual/en/function.bin2hex.php
//
// Returns an ASCII string containing the hexadecimal representation of str. The conversion is done byte-wise with the high-nibble first.
func Bin2Hex(s string) string {
	bin, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	return strconv.FormatInt(bin, 16)
}

// Explode - Split a string by a string
//
// Original : https://www.php.net/manual/en/function.explode.php
//
// Returns an array of strings, each of which is a substring of string formed by splitting it on boundaries formed by the string delimiter.
func Explode(s, set string) []string {
	return strings.Split(s, set)
}

// Hex2Bin - Decodes a hexadecimally encoded binary string
//
// Original : https://www.php.net/manual/en/function.hex2bin.php
//
// Decodes a hexadecimally encoded binary string.
//
//  NOTE: This function does NOT convert a hexadecimal number to a binary number. This can be done using the BaseConvert() function.
func Hex2Bin(s string) string {
	bin, err := strconv.ParseInt(s, 16, 64)
	if err != nil {
		log.Fatal(err)
	}
	return strconv.FormatInt(bin, 2)
}

// HtmlspecialChars - Convert special characters to HTML entities.
//
// Original : https://www.php.net/manual/en/function.htmlspecialchars.php
//
// Certain characters have special significance in HTML, and should be represented by HTML entities if they are to preserve their meanings. This function returns a string with these conversions made
func HtmlspecialChars(s string) string {
	return html.EscapeString(s)
}

// HtmlspecialCharsdecode - Convert special HTML entities back to characters.
//
// Original : https://www.php.net/manual/en/function.htmlspecialchars-decode.php
//
// Convert special HTML entities back to characters.
func HtmlspecialCharsdecode(s string) string {
	return html.UnescapeString(s)
}

// Join - Alias o Implode
//
// Original : https://www.php.net/manual/en/function.join.php
//
// Alias o Implode
func Join(set string, s []string) string {
	return Implode(set, s)
}

// Ltrim - Strip whitespace (or other characters) from the beginning of a string
//
// Original : https://www.php.net/manual/en/function.ltrim.php
//
// Strip whitespace (or other characters) from the beginning of a string.
func Ltrim(s, set string) string {
	if set == "" {
		set = " "
	}
	return strings.TrimLeft(s, set)
}

// Nl2br - Inserts HTML line breaks before all newlines in a string
//
// Original : https://www.php.net/manual/en/function.nl2br.php
//
// Returns string with <br /> or <br> inserted before all newlines (\r\n, \n\r, \n and \r).
func Nl2br(s string) string {
	return strings.ReplaceAll(s, "\n", "<br />\n")
}

// Rtrim - Strip whitespace (or other characters) from the end of a string
//
// Original : https://www.php.net/manual/en/function.rtrim.php
//
// This function returns a string with whitespace (or other characters) stripped from the end of str.
func Rtrim(s, set string) string {
	if set == "" {
		set = " "
	}
	return strings.TrimRight(s, set)
}

// StrRepeat - Repeat a string
//
// Original : https://www.php.net/manual/en/function.str-repeat.php
//
// Returns string repeated times times.
func StrRepeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// StrReplace - Replace all occurrences of the search string with the replacement string
//
// Original : https://www.php.net/manual/en/function.str-replace.php
//
// This function returns a string or an array with all occurrences of search in subject replaced with the given replace value.
//
// n value must set to -1 if you want to infinite change
func StrReplace(find, replace, s string, n int) string {
	return strings.Replace(s, find, replace, n)
}

// Strtolower - Make a string lowercase
//
// Original : https://www.php.net/manual/en/function.strtolower.php
//
// Returns string with all alphabetic characters converted to lowercase.
func Strtolower(s string) string {
	return strings.ToLower(s)
}

// Strtoupper - Make a string uppercase
//
// Original : https://www.php.net/manual/en/function.strtoupper.php
//
// Returns string with all alphabetic characters converted to uppercase.
func Strtoupper(s string) string {
	return strings.ToUpper(s)
}

// MbStrtolower - Make a string lowercase
//
// Original : https://www.php.net/manual/en/function.mb-strtolower.php
//
// Returns str with all alphabetic characters converted to lowercase.
func MbStrtolower(s string) string {
	return Strtolower(s)
}

// MbStrtoupper - Make a string uppercase
//
// Original : https://www.php.net/manual/en/function.mb-strtoupper.php
//
// Returns str with all alphabetic characters converted to uppercase.
func MbStrtoupper(s string) string {
	return Strtoupper(s)
}

// Trim - Strip whitespace (or other characters) from the beginning and end of a string
//
// Original : https://www.php.net/manual/en/function.trim.php
//
// This function returns a string with whitespace stripped from the beginning and end of str.
func Trim(s, set string) string {
	if set == "" {
		return strings.TrimSpace(s)
	}
	return strings.Trim(s, set)
}
