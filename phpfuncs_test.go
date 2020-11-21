package phpfuncs

import (
	"strings"
	"testing"
)

const TestWords = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed imperdiet iaculis libero vitae porta. Donec sit amet sodales dui. Pellentesque sit amet nibh ligula. Aliquam semper, dui ut convallis aliquet, neque diam suscipit diam, et pretium est quam quis quam. Aenean scelerisque faucibus luctus. Cras id velit sit amet turpis aliquet vehicula. Vivamus rutrum rutrum egestas."

const needle = "est"

func TestInArrayString(t *testing.T) {
	want := true
	stack := strings.Split(TestWords," ")
	if got := InArray(needle, stack); got != want {
		t.Errorf("inArray() = %v, want %v", got, want)
	}
}


func TestInArrrayInt(t *testing.T){
	want := true
	var needle = 2
	stack := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	if got := InArray(needle, stack); got != want {
		t.Errorf("inArray() = %v, want %v", got, want)
	}
}