package phpfuncs

import (
	"testing"

	"github.com/serkanalgur/phpfuncs"
)

// TestMD5 - Test for Encryption
func TestMD5(t *testing.T){
		want := "b10a8db164e0754105b7a99be72e3fe5"
		if got := phpfuncs.MD5("Hello World"); got != want {
			t.Errorf("MD5() = %v, want %v", got, want)
		}
}
