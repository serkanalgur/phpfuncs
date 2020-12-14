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

// TestBase64Decode - Test for Encryption
func TestBase64Decode(t *testing.T) {
	want := "Hello World"
	g, _ := phpfuncs.Base64Decode("SGVsbG8gV29ybGQ=")
	CoverCheck(t, g, want)
}
