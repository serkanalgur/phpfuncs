package phpfuncs

import (
	"testing"

	"github.com/serkanalgur/phpfuncs"
)

func CoverCheck(t *testing.T, s, exp string) {
	if s != exp {
		t.Errorf("%q (expected: %q)", s, exp)
	}
}

// TestBase64Encode - Test for Encryption
func TestBase64Encode(t *testing.T) {
	want := "SGVsbG8gV29ybGQ="
	CoverCheck(t,phpfuncs.Base64Encode("Hello World"),want)
}
