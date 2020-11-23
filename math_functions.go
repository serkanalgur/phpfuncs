package phpfuncs

import (
	"math"
	"strconv"
)

// Abs - Absolute value.
// Original : https://www.php.net/manual/en/function.abs.php
// Returns the absolute value of number.
func Abs(arg float64) float64 {
	return math.Abs(arg)
}

// Acos - Arc cosine.
// Original : https://www.php.net/manual/en/function.acos.php
// Returns the arc cosine of arg in radians. acos() is the inverse function of cos(), which means that a==cos(acos(a)) for every value of a that is within acos()' range.
func Acos(arg float64) float64 {
	return math.Acos(arg)
}

// Acosh -  Inverse hyperbolic cosine.
// Original : https://www.php.net/manual/en/function.acosh.php
// Returns the inverse hyperbolic cosine of arg, i.e. the value whose hyperbolic cosine is arg.
func Acosh(arg float64) float64 {
	return math.Acosh(arg)
}

// Asin - Arc sine.
// Original : https://www.php.net/manual/en/function.asin.php
// Returns the arc sine of arg in radians. asin() is the inverse function of sin(), which means that a==sin(asin(a)) for every value of a that is within asin()'s range.
func Asin(arg float64) float64 {
	return math.Asin(arg)
}

// Asinh - Inverse hyperbolic sine.
// Original : https://www.php.net/manual/en/function.asinh.php
//Returns the inverse hyperbolic sine of arg, i.e. the value whose hyperbolic sine is arg.
func Asinh(arg float64) float64 {
	return math.Asinh(arg)
}

// Atan - Arc tangent.
// Original : https://www.php.net/manual/en/function.atan.php
// Returns the arc tangent of arg in radians. atan() is the inverse function of tan(), which means that a==tan(atan(a)) for every value of a that is within atan()'s range.
func Atan(arg float64) float64 {
	return math.Atan(arg)
}

// Atan2 - Arc tangent of two variables.
// Original : https://www.php.net/manual/en/function.atan2.php
// This function calculates the arc tangent of the two variables x and y. It is similar to calculating the arc tangent of y / x, except that the signs of both arguments are used to determine the quadrant of the result.
// The function returns the result in radians, which is between -PI and PI (inclusive).
func Atan2(arg float64, arg2 float64) float64 {
	return math.Atan2(arg,arg2)
}

// Atanh - Inverse hyperbolic tangent.
// Original : https://www.php.net/manual/en/function.atanh.php
// Returns the inverse hyperbolic tangent of arg, i.e. the value whose hyperbolic tangent is arg.
func Atanh(arg float64) float64 {
	return math.Atanh(arg)
}

// BaseConvert - Convert a number between arbitrary bases.
// Original : https://www.php.net/manual/en/function.base-convert.php
// Returns a string containing number represented in base tobase. The base in which number is given is specified in frombase. Both frombase and tobase have to be between 2 and 36, inclusive. Digits in numbers with a base higher than 10 will be represented with the letters a-z, with a meaning 10, b meaning 11 and z meaning 35. The case of the letters doesn't matter, i.e. number is interpreted case-insensitively.
func BaseConvert(arg string, frombase, tobase int) (string, error) {
	ft, err := strconv.ParseInt(arg, frombase, 0)
	if err != nil {
		return "",err
	}
	return strconv.FormatInt(ft,tobase), nil
}

// Ceil - Round fractions up.
// Original : https://www.php.net/manual/en/function.ceil.php
// Returns the next highest integer value by rounding up value if necessary.
func Ceil(arg float64) float64 {
	return math.Ceil(arg)
}

// Cos - Cosine.
// Original : https://www.php.net/manual/en/function.cos.php
// cos() returns the cosine of the arg parameter. The arg parameter is in radians.
func Cos(arg float64) float64 {
	return math.Cos(arg)
}

// Cosh - Hyperbolic cosine.
// Original : https://www.php.net/manual/en/function.cosh.php
// Returns the hyperbolic cosine of arg, defined as (exp(arg) + exp(-arg))/2..
func Cosh(arg float64) float64 {
	return math.Cosh(arg)
}