/* Mathematical functions not completed
I couldn't write some of random int generators such as
getrandmax. This is a TO-DO Like comment :/
*/

package phpfuncs

import (
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// Degree180 for Deg2Rad & Rad2Deg
const Degree180 float64 = 180

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
// Returns the hyperbolic cosine of arg, defined as (exp(arg) + exp(-arg))/2.
func Cosh(arg float64) float64 {
	return math.Cosh(arg)
}

// DecBin - Hyperbolic cosine.
// Original : https://www.php.net/manual/en/function.decbin.php
// Returns a string containing a binary representation of the given number argument.
func DecBin(arg int64) string {
	return strconv.FormatInt(arg, 2)
}

// DecHex - Decimal to hexadecimal.
// Original : https://www.php.net/manual/en/function.dechex.php
// Returns a string containing a hexadecimal representation of the given unsigned number argument.
//
// The largest number that can be converted is PHP_INT_MAX * 2 + 1 (or -1): on 32-bit platforms, this will be 4294967295 in decimal, which results in dechex() returning ffffffff.
func DecHex(arg int64) string {
	return strconv.FormatInt(arg, 16)
}

// DecOct - Decimal to Octal
// Original : https://www.php.net/manual/en/function.decoct.php
// Returns a string containing an octal representation of the given number argument. The largest number that can be converted depends on the platform in use. For 32-bit platforms this is usually 4294967295 in decimal resulting in 37777777777. For 64-bit platforms this is usually 9223372036854775807 in decimal resulting in 777777777777777777777.
func DecOct(arg int64) string {
	return strconv.FormatInt(arg, 8)
}

// Deg2Rad - Converts the number in degrees to the radian equivalent.
// Original : https://www.php.net/manual/en/function.deg2rad.php
// This function converts number from degrees to the radian equivalent.
// Degree180 is constant and value is 180 :)
func Deg2Rad(arg float64) float64 {
		var mix float64 = arg / Degree180
		return mix * math.Pi
}

// Exp - Calculates the exponent of e.
// Original : https://www.php.net/manual/en/function.exp.php
// Returns e raised to the power of arg.
// Note : 'e' is the base of the natural system of logarithms, or approximately 2.718282.
func Exp(arg float64) float64 {
	return math.Exp(arg)
}

// ExpM1 - Returns exp(number) - 1, computed in a way that is accurate even when the value of number is close to zero
// Original : https://www.php.net/manual/en/function.expm1.php
// expm1() returns the equivalent to 'exp(arg) - 1' computed in a way that is accurate even if the value of arg is near zero, a case where 'exp (arg) - 1' would be inaccurate due to subtraction of two numbers that are nearly equal.
func ExpM1(arg float64) float64 {
	return math.Exp(arg) -1
}

// Floor - Round fractions down
// Original : https://www.php.net/manual/en/function.floor.php
// Returns the next lowest integer value (as float) by rounding down value if necessary.
func Floor(arg float64) float64 {
	return math.Floor(arg)
}

// FMod - Returns the floating point remainder (modulo) of the division of the arguments
// Original : https://www.php.net/manual/en/function.fmod.php
// Returns the floating point remainder of dividing the dividend (x) by the divisor (y). The remainder (r) is defined as: x = i * y + r, for some integer i. If y is non-zero, r has the same sign as x and a magnitude less than the magnitude of y.
// TO-DO : There is a problem with 0.2 like numbers. Will fix
func FMod(arg float64,arg2 float64) float64 {
		flooRit := Floor(arg / arg2)
		return arg - (flooRit * arg2)
}

// GetRandMax - Show largest possible random value
// Original : https://www.php.net/manual/en/function.getrandmax.php
// Returns the maximum value that can be returned by a call to rand().
// TO-DO
// func GetRandMax() float64{

// }

// HexDec - Hexadecimal to decimal
// Original : https://www.php.net/manual/en/function.hexdec.php
// Returns the decimal equivalent of the hexadecimal number represented by the hex_string argument. hexdec() converts a hexadecimal string to a decimal number.
//
// hexdec() will ignore any non-hexadecimal characters it encounters. As of PHP 7.4.0 supplying any invalid characters is deprecated.
func HexDec(arg string) int64 {
	decoded, err := strconv.ParseInt(arg, 16, 32)
	if err != nil {
		log.Fatal(err)
	}
	return decoded
}

// HyPot - Calculate the length of the hypotenuse of a right-angle triangle.
// Original : https://www.php.net/manual/en/function.hypot.php
// hypot() returns the length of the hypotenuse of a right-angle triangle with sides of length x and y, or the distance of the point (x, y) from the origin. This is equivalent to sqrt(x*x + y*y).
func HyPot(arg, arg2 float64) float64 {
	return math.Hypot(arg,arg2)
}

// IntDiv - Integer division.
// Original : https://www.php.net/manual/en/function.intdiv.php
// Returns the integer quotient of the division of dividend by divisor.
func IntDiv(arg, arg2 int64) (quotient int64) {
	quotient = arg / arg2
	return
}

// IsFinite - Finds whether a value is a legal finite number.
// Original : https://www.php.net/manual/en/function.is-finite.php
// Checks whether val is a legal finite on this platform.
func IsFinite(arg float64, ch int) bool {
	return !math.IsInf(arg,ch)
}

// IsInFinite - Finds whether a value is infinite.
// Original : https://www.php.net/manual/en/function.is-infinite.php
// Returns TRUE if val is infinite (positive or negative), like the result of log(0) or any value too big to fit into a float on this platform.
func IsInFinite(arg float64, ch int) bool {
	return math.IsInf(arg,ch)
}

// IsNan - Finds whether a value is infinite.
// Original : https://www.php.net/manual/en/function.is-infinite.php
// Returns TRUE if val is infinite (positive or negative), like the result of log(0) or any value too big to fit into a float on this platform.
func IsNan(arg float64) bool {
	return math.IsNaN(arg)
}

// LcgValue - Combined linear congruential generator
// Original : https://www.php.net/manual/en/function.lcg-value.php
// lcg_value() returns a pseudo random number in the range of (0, 1). The function combines two CGs with periods of 2^31 - 85 and 2^31 - 249. The period of this function is equal to the product of both primes.
func LcgValue() float64 {
	return rand.Float64()
}

// Log - Natural logarithm
// Original : https://www.php.net/manual/en/function.log.php
// If the optional base parameter is specified, log() returns logbase arg, otherwise log() returns the natural logarithm of arg.
func Log(arg float64) float64 {
	return math.Log(arg)
}

// Log10 - Base-10 logarithm
// Original : https://www.php.net/manual/en/function.log10.php
// Returns the base-10 logarithm of arg.
func Log10(arg float64) float64 {
	return math.Log10(arg)
}

// Log1P - Returns log(1 + number), computed in a way that is accurate even when the value of number is close to zero.
// Original : https://www.php.net/manual/en/function.log1p.php
// log1p() returns log(1 + number) computed in a way that is accurate even when the value of number is close to zero. log() might only return log(1) in this case due to lack of precision.
func Log1P(arg float64) float64 {
	return math.Log1p(arg)
}

// Max - Find highest value.
// Original : https://www.php.net/manual/en/function.max.php
// If at least two parameters are provided, max() returns the biggest of these values.
func Max(arg, arg2 float64) float64 {
	return math.Max(arg,arg2)
}

// Min - Find lowest value.
// Original : https://www.php.net/manual/en/function.min.php
// If at least two parameters are provided, min() returns the smallest of these values.
func Min(arg, arg2 float64) float64 {
	return math.Min(arg,arg2)
}

// OctDec - Octal to decimal
// Original : https://www.php.net/manual/en/function.octdec.php
// Returns the decimal equivalent of the octal number represented by the octal_string argument.
func OctDec(arg string) int64 {
	decoded, err := strconv.ParseInt(arg, 8, 32)
	if err != nil {
		log.Fatal(err)
	}
	return decoded
}

// Pi - Get value of pi
// Original : https://www.php.net/manual/en/function.pi.php
// Returns an approximation of pi.
func Pi() float64 {
	return math.Pi
}

// Pow - Exponential expression
// Original : https://www.php.net/manual/en/function.pow.php
// Returns base raised to the power of exp.
func Pow(arg, arg2 float64) float64 {
	return math.Pow(arg, arg2)
}

// Rad2Deg - Converts the radian number to the equivalent number in degrees.
// Original : https://www.php.net/manual/en/function.rad2deg.php
// This function converts number from radian to degrees.
// Degree180 is constant and value is 180 :)
func Rad2Deg(arg float64) float64 {
		var mix float64 = arg * Degree180
		return mix / math.Pi
}

// Rand - Generate a random integer
// Original : https://www.php.net/manual/en/function.rand.php
//  If you want a random number between 5 and 15 (inclusive), for example, use Rand(5, 15).
func Rand(args ...int) int {
	rand.Seed(time.Now().Unix())
	l := len(args)
	if l > 1 {
		min := math.Min(float64(args[0]), float64(args[1]))
		max := math.Max(float64(args[0]), float64(args[1]))
		return int(min + rand.Float64()*(max-min))
	} else if l > 0 {
		return rand.Intn(args[0])
	} else {
		return rand.Intn(1 << 32)
	}
}

// Round - Rounds a float
// Original : https://www.php.net/manual/en/function.round.php
// Returns the rounded value of val to specified precision (number of digits after the decimal point). precision can also be negative or zero (default).
func Round(arg float64) float64 {
	return math.Round(arg)
}

// Sin - Sine.
// Original : https://www.php.net/manual/en/function.sin.php
// sin() returns the sine of the arg parameter. The arg parameter is in radians.
func Sin(arg float64) float64 {
	return math.Sin(arg)
}

// Sinh - Hyperbolic sine.
// Original : https://www.php.net/manual/en/function.sinh.php
// Returns the hyperbolic sine of arg, defined as (exp(arg) - exp(-arg))/2.
func Sinh(arg float64) float64 {
	return math.Sinh(arg)
}

// Sqrt - Square root.
// Original : https://www.php.net/manual/en/function.sqrt.php
// Returns the square root of arg.
func Sqrt(arg float64) float64 {
	return math.Sqrt(arg)
}

// Tan - Tangent.
// Original : https://www.php.net/manual/en/function.tan.php
// tan() returns the tangent of the arg parameter. The arg parameter is in radians.
func Tan(arg float64) float64 {
	return math.Tan(arg)
}

// Tanh - Hyperbolic tangent.
// Original : https://www.php.net/manual/en/function.tanh.php
// Returns the hyperbolic tangent of arg, defined as sinh(arg)/cosh(arg).
func Tanh(arg float64) float64 {
	return math.Tanh(arg)
}