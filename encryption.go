package phpfuncs

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
)

// Base64Encode - Encodes data with MIME base64.
//
// Original : https://www.php.net/manual/en/function.base64-encode.php
//
// Encodes given data with base64.
func Base64Encode(v string) string {
	return base64.StdEncoding.EncodeToString([]byte(v))
}

// Base64Decode - Decodes data encoded with MIME base64.
//
// Original : https://www.php.net/manual/en/function.base64-decode.php
//
// Decodes a base64 encoded data.
func Base64Decode(v string) (string, error) {
	Base64, Err := base64.StdEncoding.DecodeString(v)
	return string(Base64), Err
}

// MD5 - Calculate the md5 hash of a string.
//
// Original : https://www.php.net/manual/en/function.md5.php
//
// Calculates the MD5 hash of str using the RSA Data Security, Inc. MD5 Message-Digest Algorithm.
func MD5(v string) string {
	hash := md5.Sum([]byte(v))
	return hex.EncodeToString(hash[:])
}

// MD5File - Calculates the md5 hash of a given file.
//
// Original : https://www.php.net/manual/en/function.md5-file.php
//
// Calculates the MD5 hash of the file specified by the filename parameter using the RSA Data Security, Inc. MD5 Message-Digest Algorithm, and returns that hash. The hash is a 32-character hexadecimal number.
func MD5File(v string) string {
	file, err := ioutil.ReadFile(v)
	if err != nil {
		return ""
	}

	hash := md5.Sum(file)
	return hex.EncodeToString(hash[:])
}

// Sha1 - Calculate the sha1 hash of a string.
//
// Original : https://www.php.net/manual/en/function.sha1.php
//
// Calculates the sha1 hash of str using the US Secure Hash Algorithm 1.
func Sha1(v string) string {
	hash := sha1.Sum([]byte(v))
	return hex.EncodeToString(hash[:])
}

// Sha224 - Calculate the sha1 hash of a string.
//
func Sha224(v string) string {
	hash := sha256.Sum224([]byte(v))
	return hex.EncodeToString(hash[:])
}

// Sha256 - Calculate the sha1 hash of a string.
//
func Sha256(v string) string {
	hash := sha256.Sum256([]byte(v))
	return hex.EncodeToString(hash[:])
}

// Sha384 - Calculate the sha1 hash of a string.
//
func Sha384(v string) string {
	hash := sha512.Sum384([]byte(v))
	return hex.EncodeToString(hash[:])
}

// Sha512 - Calculate the sha1 hash of a string.
//
func Sha512(v string) string {
	hash := sha512.Sum512([]byte(v))
	return hex.EncodeToString(hash[:])
}

// Sha1File - Calculate the sha1 hash of a file.
//
// Original : https://www.php.net/manual/en/function.sha1-file.php
//
// Calculates the sha1 hash of the file specified by filename using the US Secure Hash Algorithm 1, and returns that hash. The hash is a 40-character hexadecimal number.
func Sha1File(v string) string {
	file, err := ioutil.ReadFile(v)
	if err != nil {
		return ""
	}

	hash := sha1.Sum(file)
	return hex.EncodeToString(hash[:])
}

// Hash - Generate a hash value (message digest)
//
// Original: https://www.php.net/manual/en/function.hash.php
//
func Hash(cryp, val string) string {
	switch cryp {
	case "sha256":
		return Sha256(val)
	case "sha224":
		return Sha224(val)
	case "sha384":
		return Sha384(val)
	case "sha512":
		return Sha512(val)
	case "sha1":
		return Sha1(val)
	default:
		return MD5(val)
	}
}
