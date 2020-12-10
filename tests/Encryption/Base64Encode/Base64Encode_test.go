package phpfuncs

import (
	"testing"

	"github.com/serkanalgur/phpfuncs"
)

// TestBase64Encode - Test for Encryption
func TestBase64Encode(t *testing.T) {
	want := "SGVsbG8gV29ybGQ="
	if got := phpfuncs.Base64Encode("Hello World"); got != want {
		t.Errorf("Base64Encode() = %v, want %v", got, want)
	}
}
