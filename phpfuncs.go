package phpfuncs

import (
	"reflect"
)

/**
*	Similar function of in_array in PHP.
* Original : https://www.php.net/manual/en/function.in-array.php
* func inArray
* needle : string, int
* haystack : should be an array
* strict : set true for type check
* return boolean true / false
 */


func inArray(needle interface{}, haystack interface{}) bool {
	switch reflect.TypeOf(haystack).Kind() {
		default:
      s := reflect.ValueOf(haystack)
			for i := 0; i < s.Len(); i++ {
				if reflect.DeepEqual(needle, s.Index(i).Interface()) == true {
						return true
				}
			}
	}
	return false

}
