package phpfuncs

import "math"

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