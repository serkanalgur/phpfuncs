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

// HashCheckSha1 - Test for Encryption
func HashCheckSha1(t *testing.T) {
	want := "0a4d55a8d778e5022fab701977c5d840bbc486d0"
	CoverCheck(t, phpfuncs.Hash("sha1","Hello World"), want)
}

// HashCheckSha224 - Test for Encryption
func HashCheckSha224(t *testing.T) {
	want := "c4890faffdb0105d991a461e668e276685401b02eab1ef4372795047"
	CoverCheck(t, phpfuncs.Hash("sha224","Hello World"), want)
}

// HashCheckSha256 - Test for Encryption
func HashCheckSha256(t *testing.T) {
	want := "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e"
	CoverCheck(t, phpfuncs.Hash("sha256","Hello World"), want)
}

// HashCheckSha384 - Test for Encryption
func HashCheckSha384(t *testing.T) {
	want := "99514329186b2f6ae4a1329e7ee6c610a729636335174ac6b740f9028396fcc803d0e93863a7c3d90f86beee782f4f3f"
	CoverCheck(t, phpfuncs.Hash("sha384","Hello World"), want)
}

// HashCheckSha512 - Test for Encryption
func HashCheckSha512(t *testing.T) {
	want := "2c74fd17edafd80e8447b0d46741ee243b7eb74dd2149a0ab1b9246fb30382f27e853d8585719e0e67cbda0daa8f51671064615d645ae27acb15bfb1447f459b"
	CoverCheck(t, phpfuncs.Hash("sha512","Hello World"), want)
}

// HashCheckSha512 - Test for Encryption (Expecting to return MD5 Hash)
func HashCheckShaMD5NDefault(t *testing.T) {
	want := "b10a8db164e0754105b7a99be72e3fe5"
	CoverCheck(t, phpfuncs.Hash("sha12525","Hello World"), want)
}