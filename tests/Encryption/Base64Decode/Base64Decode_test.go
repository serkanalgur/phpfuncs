package phpfuncs

import (
	"testing"

	"github.com/serkanalgur/phpfuncs"
)

// TestBase64Decode - Test for Encryption
func TestBase64Decode(t *testing.T){
		want := "Hello World"
		if got, _ := phpfuncs.Base64Decode("SGVsbG8gV29ybGQ="); got != want {
			t.Errorf("Base64Decode() = %v, want %v", got, want)
		}
}
