package phpfuncs

import "math"

// Abs - Absolute value.
// Original : https://www.php.net/manual/en/function.abs.php
// Returns the absolute value of number.
func Abs(n float64) float64 {
	return math.Abs(n)
}

// Acos - Arc cosine.
// Original : https://www.php.net/manual/en/function.acos.php
// Returns the arc cosine of arg in radians. acos() is the inverse function of cos(), which means that a==cos(acos(a)) for every value of a that is within acos()' range.
func Acos(n float64) float64 {
	return math.Acos(n)
}

// Acosh -  Inverse hyperbolic cosine.
// Original : https://www.php.net/manual/en/function.acosh.php
// Returns the inverse hyperbolic cosine of arg, i.e. the value whose hyperbolic cosine is arg.
func Acosh(n float64) float64 {
	return math.Acosh(n)
}