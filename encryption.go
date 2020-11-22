package phpfuncs

import "encoding/base64"

// Base64Encode - Encodes data with MIME base64
// Original : https://www.php.net/manual/en/function.base64-encode.php
// Encodes given data with base64.
func Base64Encode(v string) string {
	return base64.StdEncoding.EncodeToString([]byte(v))
}

// Base64Decode - Decodes data encoded with MIME base64
// Original : https://www.php.net/manual/en/function.base64-decode.php
// Decodes a base64 encoded data.
func Base64Decode(v string) (string,error) {
	Base64, Err := base64.StdEncoding.DecodeString(v)
	return string(Base64), Err
}