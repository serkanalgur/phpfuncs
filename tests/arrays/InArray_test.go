package phpfuncs

import (
	"strings"
	"testing"

	"github.com/serkanalgur/phpfuncs"
)

const TestWords = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed imperdiet iaculis libero vitae porta. Donec sit amet sodales dui. Pellentesque sit amet nibh ligula. Aliquam semper, dui ut convallis aliquet, neque diam suscipit diam, et pretium est quam quis quam. Aenean scelerisque faucibus luctus. Cras id velit sit amet turpis aliquet vehicula. Vivamus rutrum rutrum egestas."

const needle = "est"

func CoverCheck(t *testing.T, s, exp bool) {
	if s != exp {
		t.Errorf("%v (expected: %v)", s, exp)
	}
}

func TestInArrayString(t *testing.T) {
	want := true
	stack := strings.Split(TestWords, " ")
	CoverCheck(t, phpfuncs.InArray(needle, stack), want)
}

func TestInArrrayInt(t *testing.T) {
	want := true
	var needle = 2
	stack := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	CoverCheck(t, phpfuncs.InArray(needle, stack), want)
}
