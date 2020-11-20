package phpfuncs

import (
	"strings"
	"testing"
)

const TestWords = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed imperdiet iaculis libero vitae porta. Donec sit amet sodales dui. Pellentesque sit amet nibh ligula. Aliquam semper, dui ut convallis aliquet, neque diam suscipit diam, et pretium est quam quis quam. Aenean scelerisque faucibus luctus. Cras id velit sit amet turpis aliquet vehicula. Vivamus rutrum rutrum egestas."

const needle = "est"

func TestInArray(t *testing.T) {
	want := true
	stack := strings.Split(TestWords," ")
	if got := inArray(needle, stack); got != want {
		t.Errorf("inArray() = %v, want %v", got, want)
	}
}