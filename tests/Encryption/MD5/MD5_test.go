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

// TestMD5 - Test for Encryption
func TestMD5(t *testing.T) {
	want := "b10a8db164e0754105b7a99be72e3fe5"
	CoverCheck(t, phpfuncs.MD5("Hello World"), want)
}
