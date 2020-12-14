# PHP Functions for Golang - phpfuncs

[![PkgGoDev](https://pkg.go.dev/badge/serkanalgur/phpfuncs)](https://pkg.go.dev/github.com/serkanalgur/phpfuncs) [![Go Report Card](https://goreportcard.com/badge/github.com/serkanalgur/phpfuncs)](https://goreportcard.com/report/github.com/serkanalgur/phpfuncs) [![Build Status](https://travis-ci.org/serkanalgur/phpfuncs.svg?branch=main)](https://travis-ci.org/serkanalgur/phpfuncs) [![codecov](https://codecov.io/gh/serkanalgur/phpfuncs/branch/main/graph/badge.svg?token=WETVWX4GA3)](https://codecov.io/gh/serkanalgur/phpfuncs)

PHP functions implementation to Golang. This package is for the Go beginners who have developed PHP code before. You can use PHP like functions in your app, module etc. when you add this module to your project.

:bangbang: **THIS MODULE IS UNDER DEVELOPMENT. PLEASE USE CAUTIOUSLY** :bangbang:

- [PHP Functions for Golang - phpfuncs](#php-functions-for-golang---phpfuncs)
  - [Installation](#installation)
  - [Tests](#tests)
  - [Proper Documentation](#proper-documentation)
  - [Functions List](#functions-list)
    - [func Abs](#func-abs)
    - [func Acos](#func-acos)
    - [func Acosh](#func-acosh)
    - [func Addslashes](#func-addslashes)
    - [func Array](#func-array)
    - [func ArrayCountValues](#func-arraycountvalues)
    - [func ArrayFlip](#func-arrayflip)
    - [func ArrayKeys](#func-arraykeys)
    - [func ArrayMerge](#func-arraymerge)
    - [func ArrayPush](#func-arraypush)
    - [func ArrayReverse](#func-arrayreverse)
    - [func ArrayValues](#func-arrayvalues)
    - [func Asin](#func-asin)
    - [func Asinh](#func-asinh)
    - [func Atan](#func-atan)
    - [func Atan2](#func-atan2)
    - [func Atanh](#func-atanh)
    - [func Base64Decode](#func-base64decode)
    - [func Base64Encode](#func-base64encode)
    - [func BaseConvert](#func-baseconvert)
    - [func Basename](#func-basename)
    - [func Bin2Hex](#func-bin2hex)
    - [func ByteCountIEC](#func-bytecountiec)
    - [func Ceil](#func-ceil)
    - [func Chgrp](#func-chgrp)
    - [func Chmod](#func-chmod)
    - [func Chown](#func-chown)
    - [func Copy](#func-copy)
    - [func Cos](#func-cos)
    - [func Cosh](#func-cosh)
    - [func Count](#func-count)
    - [func DecBin](#func-decbin)
    - [func DecHex](#func-dechex)
    - [func DecOct](#func-decoct)
    - [func Deg2Rad](#func-deg2rad)
    - [func Delete](#func-delete)
    - [func Die](#func-die)
    - [func DirName](#func-dirname)
    - [func Echo](#func-echo)
    - [func Exec](#func-exec)
    - [func Exit](#func-exit)
    - [func Exp](#func-exp)
    - [func ExpM1](#func-expm1)
    - [func Explode](#func-explode)
    - [func FClose](#func-fclose)
    - [func FMod](#func-fmod)
    - [func FOpen](#func-fopen)
    - [func FPuts](#func-fputs)
    - [func FRead](#func-fread)
    - [func FWrite](#func-fwrite)
    - [func FileExists](#func-fileexists)
    - [func FileGetContents](#func-filegetcontents)
    - [func FileMime](#func-filemime)
    - [func FilePerms](#func-fileperms)
    - [func FilePutContents](#func-fileputcontents)
    - [func FileSize](#func-filesize)
    - [func FileType](#func-filetype)
    - [func Floor](#func-floor)
    - [func Glob](#func-glob)
    - [func Hex2Bin](#func-hex2bin)
    - [func HexDec](#func-hexdec)
    - [func HtmlspecialChars](#func-htmlspecialchars)
    - [func HtmlspecialCharsdecode](#func-htmlspecialcharsdecode)
    - [func HyPot](#func-hypot)
    - [func Implode](#func-implode)
    - [func InArray](#func-inarray)
    - [func IntDiv](#func-intdiv)
    - [func IsArray](#func-isarray)
    - [func IsDir](#func-isdir)
    - [func IsExecutable](#func-isexecutable)
    - [func IsFile](#func-isfile)
    - [func IsFinite](#func-isfinite)
    - [func IsInFinite](#func-isinfinite)
    - [func IsLink](#func-islink)
    - [func IsNan](#func-isnan)
    - [func IsReadable](#func-isreadable)
    - [func IsURL](#func-isurl)
    - [func IsWritable](#func-iswritable)
    - [func IsWriteable](#func-iswriteable)
    - [func Join](#func-join)
    - [func LcgValue](#func-lcgvalue)
    - [func Link](#func-link)
    - [func Log](#func-log)
    - [func Log10](#func-log10)
    - [func Log1P](#func-log1p)
    - [func Ltrim](#func-ltrim)
    - [func MD5](#func-md5)
    - [func MD5File](#func-md5file)
    - [func Max](#func-max)
    - [func MbStrtolower](#func-mbstrtolower)
    - [func MbStrtoupper](#func-mbstrtoupper)
    - [func Min](#func-min)
    - [func MkDir](#func-mkdir)
    - [func Nl2br](#func-nl2br)
    - [func Now](#func-now)
    - [func OctDec](#func-octdec)
    - [func Pi](#func-pi)
    - [func Pow](#func-pow)
    - [func Print](#func-print)
    - [func PrintR](#func-printr)
    - [func Rad2Deg](#func-rad2deg)
    - [func Rand](#func-rand)
    - [func ReadLink](#func-readlink)
    - [func RealPath](#func-realpath)
    - [func Rename](#func-rename)
    - [func RmDir](#func-rmdir)
    - [func Round](#func-round)
    - [func Rtrim](#func-rtrim)
    - [func Sha1](#func-sha1)
    - [func Sha1File](#func-sha1file)
    - [func ShellExec](#func-shellexec)
    - [func Sin](#func-sin)
    - [func Sinh](#func-sinh)
    - [func Sizeof](#func-sizeof)
    - [func Sleep](#func-sleep)
    - [func Sort](#func-sort)
    - [func Sqrt](#func-sqrt)
    - [func Stat](#func-stat)
    - [func StrRepeat](#func-strrepeat)
    - [func StrReplace](#func-strreplace)
    - [func StringWithCharset](#func-stringwithcharset)
    - [func Strtolower](#func-strtolower)
    - [func Strtoupper](#func-strtoupper)
    - [func SymLink](#func-symlink)
    - [func Tan](#func-tan)
    - [func Tanh](#func-tanh)
    - [func Tempfile](#func-tempfile)
    - [func Tempnam](#func-tempnam)
    - [func Time](#func-time)
    - [func Touch](#func-touch)
    - [func Trim](#func-trim)
    - [func USleep](#func-usleep)
    - [func Unlink](#func-unlink)
    - [type ArraySlice](#type-arrayslice)
    - [func ArrayChunk](#func-arraychunk)
    - [type DiskStatus](#type-diskstatus)
    - [func DiskFreeSpace](#func-diskfreespace)

## Installation

```bash
  go get github.com/serkanalgur/phpfuncs

  # For update

  go get -u github.com/serkanalgur/phpfuncs
```

## Tests

```bash
  go test ./tests/...
```

## Proper Documentation

Please Visit [https://pkg.go.dev/github.com/serkanalgur/phpfuncs](https://pkg.go.dev/github.com/serkanalgur/phpfuncs)

## Functions List

Here is the functions list for now:

### func Abs

```go
func Abs(arg float64) float64
```

Abs - Absolute value.

Original : <https://www.php.net/manual/en/function.abs.php>

Returns the absolute value of number.

### func Acos

```go
func Acos(arg float64) float64
```

Acos - Arc cosine.

Original : <https://www.php.net/manual/en/function.acos.php>

Returns the arc cosine of arg in radians. acos() is the inverse function of
cos(), which means that a==cos(acos(a)) for every value of a that is within
acos()' range.

### func Acosh

```go
func Acosh(arg float64) float64
```

Acosh - Inverse hyperbolic cosine.

Original : <https://www.php.net/manual/en/function.acosh.php>

Returns the inverse hyperbolic cosine of arg, i.e. the value whose hyperbolic
cosine is arg.

### func Addslashes

```go
func Addslashes(s string) string
```

Addslashes - Quote string with slashes

Original : <https://www.php.net/manual/en/function.addslashes.php>

Returns a string with backslashes added before characters that need to be
escaped.

### func Array

```go
func Array(v ...interface{}) []interface{}
```

Array - Create an array

Original : <https://www.php.net/manual/en/function.array.php>

### func ArrayCountValues

```go
func ArrayCountValues(v []interface{}) map[interface{}]uint
```

ArrayCountValues - Counts all the values of an array

Original : <https://www.php.net/manual/en/function.array-count-values.php>

### func ArrayFlip

```go
func ArrayFlip(v map[interface{}]interface{}) map[interface{}]interface{}
```

ArrayFlip - Exchanges all keys with their associated values in an array

Original : <https://www.php.net/manual/en/function.array-flip.php>

ArrayFlip returns an array in flip order, i.e. keys from array become values and
values from array become keys.

### func ArrayKeys

```go
func ArrayKeys(v map[string]interface{}) []string
```

ArrayKeys - Return all the keys or a subset of the keys of an array

Original : <https://www.php.net/manual/en/function.array-keys.php>

### func ArrayMerge

```go
func ArrayMerge(v ...[]interface{}) []interface{}
```

ArrayMerge - Merge one or more arrays

Original : <https://www.php.net/manual/en/function.array-merge.php>

### func ArrayPush

```go
func ArrayPush(v *[]interface{}, data ...interface{})
```

ArrayPush - Push one or more elements onto the end of array

Original : <https://www.php.net/manual/en/function.array-push.php>

### func ArrayReverse

```go
func ArrayReverse(v []interface{}) []interface{}
```

ArrayReverse - Return an array with elements in reverse order

Original : <https://www.php.net/manual/en/function.array-reverse.php>

### func ArrayValues

```go
func ArrayValues(v ...[]interface{}) (value []interface{})
```

ArrayValues - Return all the values of an array

Original : <https://www.php.net/manual/en/function.array-values.php>

### func Asin

```go
func Asin(arg float64) float64
```

Asin - Arc sine.

Original : <https://www.php.net/manual/en/function.asin.php>

Returns the arc sine of arg in radians. asin() is the inverse function of sin(),
which means that a==sin(asin(a)) for every value of a that is within asin()'s
range.

### func Asinh

```go
func Asinh(arg float64) float64
```

Asinh - Inverse hyperbolic sine.

Original : <https://www.php.net/manual/en/function.asinh.php>

Returns the inverse hyperbolic sine of arg, i.e. the value whose hyperbolic sine
is arg.

### func Atan

```go
func Atan(arg float64) float64
```

Atan - Arc tangent.

Original : <https://www.php.net/manual/en/function.atan.php>

Returns the arc tangent of arg in radians. atan() is the inverse function of
tan(), which means that a==tan(atan(a)) for every value of a that is within
atan()'s range.

### func Atan2

```go
func Atan2(arg float64, arg2 float64) float64
```

Atan2 - Arc tangent of two variables.

Original : <https://www.php.net/manual/en/function.atan2.php>

This function calculates the arc tangent of the two variables x and y. It is
similar to calculating the arc tangent of y / x, except that the signs of both
arguments are used to determine the quadrant of the result.

The function returns the result in radians, which is between -PI and PI
(inclusive).

### func Atanh

```go
func Atanh(arg float64) float64
```

Atanh - Inverse hyperbolic tangent.

Original : <https://www.php.net/manual/en/function.atanh.php>

Returns the inverse hyperbolic tangent of arg, i.e. the value whose hyperbolic
tangent is arg.

### func Base64Decode

```go
func Base64Decode(v string) (string, error)
```

Base64Decode - Decodes data encoded with MIME base64.

Original : <https://www.php.net/manual/en/function.base64-decode.php>

Decodes a base64 encoded data.

### func Base64Encode

```go
func Base64Encode(v string) string
```

Base64Encode - Encodes data with MIME base64.

Original : <https://www.php.net/manual/en/function.base64-encode.php>

Encodes given data with base64.

### func BaseConvert

```go
func BaseConvert(arg string, frombase, tobase int) (string, error)
```

BaseConvert - Convert a number between arbitrary bases.

Original : <https://www.php.net/manual/en/function.base-convert.php>

Returns a string containing number represented in base tobase. The base in which
number is given is specified in frombase. Both frombase and tobase have to be
between 2 and 36, inclusive. Digits in numbers with a base higher than 10 will
be represented with the letters a-z, with a meaning 10, b meaning 11 and z
meaning 35. The case of the letters doesn't matter, i.e. number is interpreted
case-insensitively.

### func Basename

```go
func Basename(path string) string
```

Basename - Returns trailing name component of path.

Original : <https://www.php.net/manual/en/function.basename.php>

Given a string containing the path to a file or directory, this function will
return the trailing name component.

### func Bin2Hex

```go
func Bin2Hex(s string) string
```

Bin2Hex - Convert binary data into hexadecimal representation

Original : <https://www.php.net/manual/en/function.bin2hex.php>

Returns an ASCII string containing the hexadecimal representation of str. The
conversion is done byte-wise with the high-nibble first.

### func ByteCountIEC

```go
func ByteCountIEC(b uint64) string
```

ByteCountIEC - Bytecount & Humanize Bytes

Complete calculator for DiskFreeSize

### func Ceil

```go
func Ceil(arg float64) float64
```

Ceil - Round fractions up.

Original : <https://www.php.net/manual/en/function.ceil.php>

Returns the next highest integer value by rounding up value if necessary.

### func Chgrp

```go
func Chgrp(name string, uid, gid int) error
```

Chgrp - Changes file group.

Original : <https://www.php.net/manual/en/function.chgrp.php>

Attempts to change the group of the file filename to group.

Only the superuser may change the group of a file arbitrarily; other users may
change the group of a file to any group of which that user is a member.

### func Chmod

```go
func Chmod(name string, mode os.FileMode) error
```

Chmod - Changes file mode

Original : <https://www.php.net/manual/en/function.chmod.php>

Attempts to change the mode of the specified file to that given in mode.

### func Chown

```go
func Chown(name string, uid int, gid int) error
```

Chown - Changes file owner.

Original : <https://www.php.net/manual/en/function.chown.php>

Attempts to change the owner of the file filename to user user. Only the
superuser may change the owner of a file.

### func Copy

```go
func Copy(src, dst string) (int64, error)
```

Copy - Copies file

Original : <https://www.php.net/manual/en/function.copy.php>

Makes a copy of the file source to dest.

### func Cos

```go
func Cos(arg float64) float64
```

Cos - Cosine.

Original : <https://www.php.net/manual/en/function.cos.php>

cos() returns the cosine of the arg parameter. The arg parameter is in radians.

### func Cosh

```go
func Cosh(arg float64) float64
```

Cosh - Hyperbolic cosine.

Original : <https://www.php.net/manual/en/function.cosh.php>

Returns the hyperbolic cosine of arg, defined as (exp(arg) + exp(-arg))/2.

### func Count

```go
func Count(v []interface{}) int
```

Count - Count all elements in an array, or something in an object

Original : <https://www.php.net/manual/en/function.count.php>

### func DecBin

```go
func DecBin(arg int64) string
```

DecBin - Hyperbolic cosine.

Original : <https://www.php.net/manual/en/function.decbin.php>

Returns a string containing a binary representation of the given number
argument.

### func DecHex

```go
func DecHex(arg int64) string
```

DecHex - Decimal to hexadecimal.

Original : <https://www.php.net/manual/en/function.dechex.php>

Returns a string containing a hexadecimal representation of the given unsigned
number argument.

The largest number that can be converted is PHP_INT_MAX \* 2 + 1 (or -1): on
32-bit platforms, this will be 4294967295 in decimal, which results in dechex()
returning ffffffff.

### func DecOct

```go
func DecOct(arg int64) string
```

DecOct - Decimal to Octal

Original : <https://www.php.net/manual/en/function.decoct.php>

Returns a string containing an octal representation of the given number
argument. The largest number that can be converted depends on the platform in
use. For 32-bit platforms this is usually 4294967295 in decimal resulting in 37777777777. For 64-bit platforms this is usually 9223372036854775807 in decimal
resulting in 777777777777777777777.

### func Deg2Rad

```go
func Deg2Rad(arg float64) float64
```

Deg2Rad - Converts the number in degrees to the radian equivalent.

Original : <https://www.php.net/manual/en/function.deg2rad.php>

This function converts number from degrees to the radian equivalent.

Degree180 is constant and value is 180 :)

### func Delete

```go
func Delete(name string) error
```

Delete - Deletes a file.

Original : <https://www.php.net/manual/en/function.delete.php>

Deletes filename. Similar to the Unix C unlink() function. An E_WARNING level
error will be generated on failure.

### func Die

```go
func Die(of int)
```

Die - Equivalent to exit

Original : <https://www.php.net/manual/en/function.die.php>

This language construct is equivalent to exit().

### func DirName

```go
func DirName(path string) ([]os.FileInfo, error)
```

DirName - Returns a parent directory's path

Original : <https://www.php.net/manual/en/function.dirname.php>

Given a string containing the path of a file or directory, this function will
return the parent directory's path that is levels up from the current directory.

### func Echo

```go
func Echo(vars ...interface{})
```

Echo - Output one or more strings

Original : <https://www.php.net/manual/en/function.echo.php>

### func Exec

```go
func Exec(of string)
```

Exec - Start a command on system

Original : <https://www.php.net/manual/tr/function.exec.php>

exec() executes the given command.

### func Exit

```go
func Exit(of int)
```

Exit - Output a message and terminate the current script

Original : <https://www.php.net/manual/en/function.exit.php>

Terminates execution of the script. Shutdown functions and object destructors
will always be executed even if exit is called.

### func Exp

```go
func Exp(arg float64) float64
```

Exp - Calculates the exponent of e.

Original : <https://www.php.net/manual/en/function.exp.php>

Returns e raised to the power of arg.

Note : 'e' is the base of the natural system of logarithms, or approximately
2.718282.

### func ExpM1

```go
func ExpM1(arg float64) float64
```

ExpM1 - Returns exp(number) - 1, computed in a way that is accurate even when
the value of number is close to zero

Original : <https://www.php.net/manual/en/function.expm1.php>

expm1() returns the equivalent to 'exp(arg) - 1' computed in a way that is
accurate even if the value of arg is near zero, a case where 'exp (arg) - 1'
would be inaccurate due to subtraction of two numbers that are nearly equal.

### func Explode

```go
func Explode(s, set string) []string
```

Explode - Split a string by a string

Original : <https://www.php.net/manual/en/function.explode.php>

Returns an array of strings, each of which is a substring of string formed by
splitting it on boundaries formed by the string delimiter.

### func FClose

```go
func FClose(file *os.File) error
```

FClose - Closes an open file pointer

Original : <https://www.php.net/manual/en/function.fclose.php>

The file pointed to by handle is closed.

### func FMod

```go
func FMod(arg float64, arg2 float64) float64
```

FMod - Returns the floating point remainder (modulo) of the division of the
arguments

Original : <https://www.php.net/manual/en/function.fmod.php>

Returns the floating point remainder of dividing the dividend (x) by the divisor
(y). The remainder (r) is defined as: x = i \* y + r, for some integer i. If y is
non-zero, r has the same sign as x and a magnitude less than the magnitude of y.

TO-DO : There is a problem with 0.2 like numbers. Will fix

### func FOpen

```go
func FOpen(file string, mode int) (*os.File, error)
```

FOpen - Opens file

Original : <https://www.php.net/manual/en/function.fopen.php>

fopen() binds a named resource, specified by filename, to a stream.

Mode : os.O_RDONLY | os.O_WRONLY | os.O_RDWR | os.O_APPEND | os.O_CREATE |
os.O_EXCL | os.O_SYNC | os.O_TRUNC

### func FPuts

```go
func FPuts(f *os.File, data string) int
```

FPuts - Alias of FWrite

Original : <https://www.php.net/manual/en/function.fputs.php>

This function is an alias of: FWrite()..

### func FRead

```go
func FRead(f *os.File, sb int64) string
```

FRead - Binary-safe file read.

Original : <https://www.php.net/manual/en/function.fread.php>

fread() reads up to length bytes from the file pointer referenced by handle.

### func FWrite

```go
func FWrite(f *os.File, data string) int
```

FWrite - Binary-safe file write

Original : <https://www.php.net/manual/en/function.fwrite.php>

fwrite() writes the contents of string to the file stream pointed to by handle.

### func FileExists

```go
func FileExists(path string) bool
```

FileExists - Checks whether a file or directory exists.

Original : <https://www.php.net/manual/en/function.file-exists.php>

Checks whether a file or directory exists.

### func FileGetContents

```go
func FileGetContents(path string, includePath bool, context []string, offset int, maxlen int) string
```

FileGetContents - Reads entire file into a string.

Original : <https://www.php.net/manual/en/function.file-get-contents.php>

This function is similar to file(), except that file_get_contents() returns the
file in a string, starting at the specified offset up to maxlen bytes. On
failure, file_get_contents() will return FALSE. TODO : Context Implementation.

### func FileMime

```go
func FileMime(file string) time.Time
```

FileMime - Gets file modification time

Original : <https://www.php.net/manual/en/function.filemtime.php>

This function returns the time when the data blocks of a file were being written
to, that is, the time when the content of the file was changed.

### func FilePerms

```go
func FilePerms(path string) os.FileMode
```

FilePerms - Gets file permissions.

Original : <https://www.php.net/manual/en/function.fileperms.php> Gets permissions
for the given file.

### func FilePutContents

```go
func FilePutContents(path, data string)
```

FilePutContents - Write data to a file

Original : <https://www.php.net/manual/en/function.file-put-contents.php>

This function is identical to calling fopen(), fwrite() and fclose()
successively to write data to a file. TODO: Flags

### func FileSize

```go
func FileSize(path string) (int64, error)
```

FileSize - Gets file permissions.

Original : <https://www.php.net/manual/en/function.filesize.php> Gets permissions
for the given file.

### func FileType

```go
func FileType(fs string) (string, error)
```

FileType - Gets file type.

Original : <https://www.php.net/manual/en/function.filetype.php>

Returns the type of the given file.

### func Floor

```go
func Floor(arg float64) float64
```

Floor - Round fractions down

Original : <https://www.php.net/manual/en/function.floor.php>

Returns the next lowest integer value (as float) by rounding down value if
necessary.

### func Glob

```go
func Glob(path string) (matches []string, err error)
```

Glob - Find pathnames matching a pattern.

Original : <https://www.php.net/manual/en/function.glob.php>

The glob() function searches for all the pathnames matching pattern according to
the rules used by the libc glob() function, which is similar to the rules used
by common shells.

### func Hex2Bin

```go
func Hex2Bin(s string) string
```

Hex2Bin - Decodes a hexadecimally encoded binary string

Original : <https://www.php.net/manual/en/function.hex2bin.php>

Decodes a hexadecimally encoded binary string.

```NOTE: This function does NOT convert a hexadecimal number to a binary number. This can be done using the BaseConvert() function.```

### func HexDec

```go
func HexDec(arg string) int64
```

HexDec - Hexadecimal to decimal

Original : <https://www.php.net/manual/en/function.hexdec.php>

Returns the decimal equivalent of the hexadecimal number represented by the
hex_string argument. hexdec() converts a hexadecimal string to a decimal number.

hexdec() will ignore any non-hexadecimal characters it encounters. As of PHP
7.4.0 supplying any invalid characters is deprecated.

### func HtmlspecialChars

```go
func HtmlspecialChars(s string) string
```

HtmlspecialChars - Convert special characters to HTML entities.

Original : <https://www.php.net/manual/en/function.htmlspecialchars.php>

Certain characters have special significance in HTML, and should be represented
by HTML entities if they are to preserve their meanings. This function returns a
string with these conversions made

### func HtmlspecialCharsdecode

```go
func HtmlspecialCharsdecode(s string) string
```

HtmlspecialCharsdecode - Convert special HTML entities back to characters.

Original : <https://www.php.net/manual/en/function.htmlspecialchars-decode.php>

Convert special HTML entities back to characters.

### func HyPot

```go
func HyPot(arg, arg2 float64) float64
```

HyPot - Calculate the length of the hypotenuse of a right-angle triangle.

Original : <https://www.php.net/manual/en/function.hypot.php>

hypot() returns the length of the hypotenuse of a right-angle triangle with
sides of length x and y, or the distance of the point (x, y) from the origin.
This is equivalent to sqrt(x*x + y*y).

### func Implode

```go
func Implode(sep string, v []string) string
```

Implode - Join array elements with a string.

Original : <https://www.php.net/manual/en/function.implode.php>

Join array elements with a glue string.

### func InArray

```go
func InArray(needle interface{}, haystack interface{}) bool
```

InArray - Similar function of in_array in PHP

Original : <https://www.php.net/manual/en/function.in-array.php>

  ```needle : string, int
    haystack : should be an array
    strict : set true for type check
    return boolean true / false
  ```

### func IntDiv

```go
func IntDiv(arg, arg2 int64) (quotient int64)
```

IntDiv - Integer division.

Original : <https://www.php.net/manual/en/function.intdiv.php>

Returns the integer quotient of the division of dividend by divisor.

### func IsArray

```go
func IsArray(v interface{}) bool
```

IsArray - Finds whether a variable is an array.

Original : <https://www.php.net/manual/en/function.is-array.php>

Finds whether the given variable is an array.

### func IsDir

```go
func IsDir(path string) bool
```

IsDir - Tells whether the filename is a directory.

Original : <https://www.php.net/manual/en/function.is-dir.php>

Tells whether the given filename is a directory.

### func IsExecutable

```go
func IsExecutable(path string) bool
```

IsExecutable - Tells whether the filename is executable

Original : <https://www.php.net/manual/en/function.is-executable.php>

Tells whether the filename is executable.

### func IsFile

```go
func IsFile(path string) bool
```

IsFile - Tells whether the filename is a regular file.

Original : <https://www.php.net/manual/en/function.is-file.php>

Tells whether the given file is a regular file.

### func IsFinite

```go
func IsFinite(arg float64, ch int) bool
```

IsFinite - Finds whether a value is a legal finite number.

Original : <https://www.php.net/manual/en/function.is-finite.php>

Checks whether val is a legal finite on this platform.

### func IsInFinite

```go
func IsInFinite(arg float64, ch int) bool
```

IsInFinite - Finds whether a value is infinite.

Original : <https://www.php.net/manual/en/function.is-infinite.php>

Returns TRUE if val is infinite (positive or negative), like the result of
log(0) or any value too big to fit into a float on this platform.

### func IsLink

```go
func IsLink(path string) bool
```

IsLink - Tells whether the filename is a symbolic link.

Original : <https://www.php.net/manual/en/function.is-link.php>

Tells whether the given file is a symbolic link.

### func IsNan

```go
func IsNan(arg float64) bool
```

IsNan - Finds whether a value is infinite.

Original : <https://www.php.net/manual/en/function.is-infinite.php>

Returns TRUE if val is infinite (positive or negative), like the result of
log(0) or any value too big to fit into a float on this platform.

### func IsReadable

```go
func IsReadable(path string) bool
```

IsReadable - Tells whether a file exists and is readable.

Original : <https://www.php.net/manual/en/function.is-readable.php>

Tells whether a file exists and is readable.

### func IsURL

```go
func IsURL(str string) bool
```

IsURL function

### func IsWritable

```go
func IsWritable(path string) bool
```

IsWritable - Tells whether the filename is writable.

Original : <https://www.php.net/manual/en/function.is-writable.php>

Returns TRUE if the filename exists and is writable. The filename argument may
be a directory name allowing you to check if a directory is writable.

### func IsWriteable

```go
func IsWriteable(path string) bool
```

IsWriteable - Tells whether the filename is writable.

Original : <https://www.php.net/manual/en/function.is-writeable.php>

Returns TRUE if the filename exists and is writable. The filename argument may
be a directory name allowing you to check if a directory is writable.

### func Join

```go
func Join(set string, s []string) string
```

Join - Alias o Implode

Original : <https://www.php.net/manual/en/function.join.php>

Alias o Implode

### func LcgValue

```go
func LcgValue() float64
```

LcgValue - Combined linear congruential generator

Original : <https://www.php.net/manual/en/function.lcg-value.php>

lcg_value() returns a pseudo random number in the range of (0, 1). The function
combines two CGs with periods of 2^31 - 85 and 2^31 - 249. The period of this
function is equal to the product of both primes.

### func Link

```go
func Link(target, link string)
```

Link - Create a hard link

Original : <https://www.php.net/manual/en/function.link.php>

link() creates a hard link.

### func Log

```go
func Log(arg float64) float64
```

Log - Natural logarithm

Original : <https://www.php.net/manual/en/function.log.php>

If the optional base parameter is specified, log() returns logbase arg,
otherwise log() returns the natural logarithm of arg.

### func Log10

```go
func Log10(arg float64) float64
```

Log10 - Base-10 logarithm

Original : <https://www.php.net/manual/en/function.log10.php>

Returns the base-10 logarithm of arg.

### func Log1P

```go
func Log1P(arg float64) float64
```

Log1P - Returns log(1 + number), computed in a way that is accurate even when
the value of number is close to zero.

Original : <https://www.php.net/manual/en/function.log1p.php>

log1p() returns log(1 + number) computed in a way that is accurate even when the
value of number is close to zero. log() might only return log(1) in this case
due to lack of precision.

### func Ltrim

```go
func Ltrim(s, set string) string
```

Ltrim - Strip whitespace (or other characters) from the beginning of a string

Original : <https://www.php.net/manual/en/function.ltrim.php>

Strip whitespace (or other characters) from the beginning of a string.

### func MD5

```go
func MD5(v string) string
```

MD5 - Calculate the md5 hash of a string.

Original : <https://www.php.net/manual/en/function.md5.php>

Calculates the MD5 hash of str using the RSA Data Security, Inc. MD5
Message-Digest Algorithm.

### func MD5File

```go
func MD5File(v string) string
```

MD5File - Calculates the md5 hash of a given file.

Original : <https://www.php.net/manual/en/function.md5-file.php>

Calculates the MD5 hash of the file specified by the filename parameter using
the RSA Data Security, Inc. MD5 Message-Digest Algorithm, and returns that hash.
The hash is a 32-character hexadecimal number.

### func Max

```go
func Max(arg, arg2 float64) float64
```

Max - Find highest value.

Original : <https://www.php.net/manual/en/function.max.php>

If at least two parameters are provided, max() returns the biggest of these
values.

### func MbStrtolower

```go
func MbStrtolower(s string) string
```

MbStrtolower - Make a string lowercase

Original : <https://www.php.net/manual/en/function.mb-strtolower.php>

Returns str with all alphabetic characters converted to lowercase.

### func MbStrtoupper

```go
func MbStrtoupper(s string) string
```

MbStrtoupper - Make a string uppercase

Original : <https://www.php.net/manual/en/function.mb-strtoupper.php>

Returns str with all alphabetic characters converted to uppercase.

### func Min

```go
func Min(arg, arg2 float64) float64
```

Min - Find lowest value.

Original : <https://www.php.net/manual/en/function.min.php>

If at least two parameters are provided, min() returns the smallest of these
values.

### func MkDir

```go
func MkDir(path string, mode os.FileMode) error
```

MkDir - Makes directory.

Original : <https://www.php.net/manual/en/function.mkdir.php>

Attempts to create the directory specified by pathname.

### func Nl2br

```go
func Nl2br(s string) string
```

Nl2br - Inserts HTML line breaks before all newlines in a string

Original : <https://www.php.net/manual/en/function.nl2br.php>

Returns string with ```<br />``` or ```<br>``` inserted before all newlines (\r\n, \n\r, \n
and \r).

### func Now

```go
func Now() int64
```

Now - Similar function of time() in PHP.

Original : <https://www.php.net/manual/en/function.time.php>

Returns the current time measured in the number of seconds since the Unix Epoch
(January 1 1970 00:00:00 GMT).

### func OctDec

```go
func OctDec(arg string) int64
```

OctDec - Octal to decimal

Original : <https://www.php.net/manual/en/function.octdec.php>

Returns the decimal equivalent of the octal number represented by the
octal_string argument.

### func Pi

```go
func Pi() float64
```

Pi - Get value of pi

Original : <https://www.php.net/manual/en/function.pi.php>

Returns an approximation of pi.

### func Pow

```go
func Pow(arg, arg2 float64) float64
```

Pow - Exponential expression

Original : <https://www.php.net/manual/en/function.pow.php>

Returns base raised to the power of exp.

### func Print

```go
func Print(vars string)
```

Print - Output a string

Original : <https://www.php.net/manual/en/function.print.php>

### func PrintR

```go
func PrintR(vars ...interface{})
```

PrintR - Prints human-readable information about a variable

Original : <https://www.php.net/manual/en/function.print-r.php>

### func Rad2Deg

```go
func Rad2Deg(arg float64) float64
```

Rad2Deg - Converts the radian number to the equivalent number in degrees.

Original : <https://www.php.net/manual/en/function.rad2deg.php>

This function converts number from radian to degrees. Degree180 is constant and
value is 180 :)

### func Rand

```go
func Rand(args ...int) int
```

Rand - Generate a random integer

Original : <https://www.php.net/manual/en/function.rand.php>

If you want a random number between 5 and 15 (inclusive), for example, use
Rand(5, 15).

### func ReadLink

```go
func ReadLink(path string) (string, error)
```

ReadLink - Returns the target of a symbolic link.

Original : <https://www.php.net/manual/en/function.readlink.php>

readlink() does the same as the readlink C function.

### func RealPath

```go
func RealPath(path string) (string, error)
```

RealPath - Returns canonicalized absolute pathname.

Original : <https://www.php.net/manual/en/function.realpath.php>

realpath() expands all symbolic links and resolves references to /./, /../ and
extra / characters in the input path and returns the canonicalized absolute
pathname.

### func Rename

```go
func Rename(oldpath, newpath string) error
```

Rename - Renames a file or directory.

Original : <https://www.php.net/manual/en/function.rename.php>

Attempts to rename oldname to newname, moving it between directories if
necessary. If renaming a file and newname exists, it will be overwritten. If
renaming a directory and newname exists, this function will emit a warning.

### func RmDir

```go
func RmDir(path string) error
```

RmDir — Removes directory.

Original : <https://www.php.net/manual/en/function.rmdir.php>

Attempts to remove the directory named by dirname. The directory must be empty,
and the relevant permissions must permit this. A E_WARNING level error will be
generated on failure.

### func Round

```go
func Round(arg float64) float64
```

Round - Rounds a float

Original : <https://www.php.net/manual/en/function.round.php>

Returns the rounded value of val to specified precision (number of digits after
the decimal point). precision can also be negative or zero (default).

### func Rtrim

```go
func Rtrim(s, set string) string
```

Rtrim - Strip whitespace (or other characters) from the end of a string

Original : <https://www.php.net/manual/en/function.rtrim.php>

This function returns a string with whitespace (or other characters) stripped
from the end of str.

### func Sha1

```go
func Sha1(v string) string
```

Sha1 - Calculate the sha1 hash of a string.

Original : <https://www.php.net/manual/en/function.sha1.php>

Calculates the sha1 hash of str using the US Secure Hash Algorithm 1.

### func Sha1File

```go
func Sha1File(v string) string
```

Sha1File - Calculate the sha1 hash of a file.

Original : <https://www.php.net/manual/en/function.sha1-file.php>

Calculates the sha1 hash of the file specified by filename using the US Secure
Hash Algorithm 1, and returns that hash. The hash is a 40-character hexadecimal
number.

### func ShellExec

```go
func ShellExec(of string)
```

ShellExec - Execute command via shell and return the complete output as a string

Original : <https://www.php.net/manual/en/function.shell-exec.php>

This function is identical to the backtick operator.

### func Sin

```go
func Sin(arg float64) float64
```

Sin - Sine.

Original : <https://www.php.net/manual/en/function.sin.php>

sin() returns the sine of the arg parameter. The arg parameter is in radians.

### func Sinh

```go
func Sinh(arg float64) float64
```

Sinh - Hyperbolic sine.

Original : <https://www.php.net/manual/en/function.sinh.php>

Returns the hyperbolic sine of arg, defined as (exp(arg) - exp(-arg))/2.

### func Sizeof

```go
func Sizeof(v []interface{}) int
```

Sizeof - Alias of Count

Original : <https://www.php.net/manual/en/function.sizeof.php>

### func Sleep

```go
func Sleep(t int64)
```

Sleep - Delay in seconds

Original : <https://www.php.net/manual/en/function.sleep.php>

Delays the program execution for the given number of seconds.

### func Sort

```go
func Sort(v interface{}, flag string) bool
```

Sort - Sort an array

Original : <https://www.php.net/manual/en/function.sort.php>

This function sorts an array. Elements will be arranged from lowest to highest
when this function has completed.

Flags are same as php sort.

  ```SORT_REGULAR - compare items normally;
  SORT_NUMERIC - compare items numerically
  SORT_STRING - compare items as strings
  SORT_NATURAL - compare items as strings using "natural ordering" like natsort()
  ```

### func Sqrt

```go
func Sqrt(arg float64) float64
```

Sqrt - Square root.

Original : <https://www.php.net/manual/en/function.sqrt.php>

Returns the square root of arg.

### func Stat

```go
func Stat(name string) (os.FileInfo, error)
```

Stat - Gives information about a file.

Original : <https://www.php.net/manual/en/function.stat.php>

Gathers the statistics of the file named by filename. If filename is a symbolic
link, statistics are from the file itself, not the symlink. Prior to PHP 7.4.0,
on Windows NTS builds the size, atime, mtime and ctime statistics have been from
the symlink, in this case.

### func StrRepeat

```go
func StrRepeat(s string, count int) string
```

StrRepeat - Repeat a string

Original : <https://www.php.net/manual/en/function.str-repeat.php>

Returns string repeated times times.

### func StrReplace

```go
func StrReplace(find, replace, s string, n int) string
```

StrReplace - Replace all occurrences of the search string with the replacement
string

Original : <https://www.php.net/manual/en/function.str-replace.php>

This function returns a string or an array with all occurrences of search in
subject replaced with the given replace value.

n value must set to -1 if you want to infinite change

### func StringWithCharset

```go
func StringWithCharset(length int, charset string) string
```

StringWithCharset - Create RAndom and Unique String

### func Strtolower

```go
func Strtolower(s string) string
```

Strtolower - Make a string lowercase

Original : <https://www.php.net/manual/en/function.strtolower.php>

Returns string with all alphabetic characters converted to lowercase.

### func Strtoupper

```go
func Strtoupper(s string) string
```

Strtoupper - Make a string uppercase

Original : <https://www.php.net/manual/en/function.strtoupper.php>

Returns string with all alphabetic characters converted to uppercase.

### func SymLink

```go
func SymLink(target, link string)
```

SymLink - Creates a symbolic link

Original : <https://www.php.net/manual/en/function.symlink.php>

symlink() creates a symbolic link to the existing target with the specified name
link.

### func Tan

```go
func Tan(arg float64) float64
```

Tan - Tangent.

Original : <https://www.php.net/manual/en/function.tan.php>

tan() returns the tangent of the arg parameter. The arg parameter is in radians.

### func Tanh

```go
func Tanh(arg float64) float64
```

Tanh - Hyperbolic tangent.

Original : <https://www.php.net/manual/en/function.tanh.php>

Returns the hyperbolic tangent of arg, defined as sinh(arg)/cosh(arg).

### func Tempfile

```go
func Tempfile() (f *os.File)
```

Tempfile - Creates a temporary file

Original : <https://www.php.net/manual/en/function.tmpfile.php>

Creates a temporary file with a unique name in read-write (w+) mode and returns
a file handle.

The file is automatically removed when closed (for example, by calling fclose(),
or when there are no remaining references to the file handle returned by
tmpfile()), or when the script ends.

### func Tempnam

```go
func Tempnam(dir, prefix string) string
```

Tempnam - Create file with unique file name

Original : <https://www.php.net/manual/en/function.tempnam.php>

Creates a file with a unique filename, with access permission set to 0600, in
the specified directory. If the directory does not exist or is not writable,
tempnam() may generate a file in the system's temporary directory, and return
the full path to that file, including its name.

### func Time

```go
func Time() int64
```

Time - Similar function of time() in PHP.

Original : <https://www.php.net/manual/en/function.time.php>

Returns the current time measured in the number of seconds since the Unix Epoch
(January 1 1970 00:00:00 GMT).

### func Touch

```go
func Touch(path string, t int64, at int64) bool
```

Touch - Sets access and modification time of file

Original : <https://www.php.net/manual/en/function.touch.php>

Attempts to set the access and modification times of the file named in the
filename parameter to the value given in time. Note that the access time is
always modified, regardless of the number of parameters. If the file does not
exist, it will be created.

### func Trim

```go
func Trim(s, set string) string
```

Trim - Strip whitespace (or other characters) from the beginning and end of a
string

Original : <https://www.php.net/manual/en/function.trim.php>

This function returns a string with whitespace stripped from the beginning and
end of str.

### func USleep

```go
func USleep(t int64)
```

USleep - Delay in microseconds

Original : <https://www.php.net/manual/en/function.usleep.php>

Delays program execution for the given number of microseconds.

### func Unlink

```go
func Unlink(name string) error
```

Unlink - Deletes a file.

Original : <https://www.php.net/manual/en/function.unlink.php>

Deletes filename. Similar to the Unix C unlink() function. An E_WARNING level
error will be generated on failure.

### type ArraySlice

```go
type ArraySlice []interface{}
```

ArraySlice type -- Required for ArrayChunk

### func ArrayChunk

```go
func ArrayChunk(v ArraySlice, size int) ArraySlice
```

ArrayChunk - Split an array into chunks

Original : <https://www.php.net/manual/en/function.array-chunk.php>

### type DiskStatus

```go
type DiskStatus struct {
 Free string `json:"Free"`
}
```

DiskStatus struct

### func DiskFreeSpace

```go
func DiskFreeSpace(path string) (disk DiskStatus)
```

DiskFreeSpace - Returns available space on filesystem or disk partition

Original : <https://www.php.net/manual/en/function.disk-free-space.php>

Given a string containing a directory, this function will return the number of
bytes available on the corresponding filesystem or disk partition.

DEVELOPER NOTE : PROBABLY WORKING ON ONLY LINUX AND MAC. TO-DO : WINDOWS
