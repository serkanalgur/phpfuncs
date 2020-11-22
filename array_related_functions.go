package phpfuncs

import (
	"fmt"
	"math"
)

// ArraySlice type -- Required for ArrayChunk
type ArraySlice []interface{}

// Array - Create an array
// Original : https://www.php.net/manual/en/function.array.php
func Array(v... interface{}) []interface{} {
	return v
}

// ArrayKeys - Return all the keys or a subset of the keys of an array
// Original : https://www.php.net/manual/en/function.array-keys.php
func ArrayKeys(v map[string]interface{}) []string {
	var aRdata []string

	for s := range v {
		aRdata = append(aRdata, s)
	}

	return aRdata
}

// ArrayPush - Push one or more elements onto the end of array
// Original : https://www.php.net/manual/en/function.array-push.php
func ArrayPush(v *[]interface{}, data ...interface{}) {
		*v = append(*v, data...)
}

// ArrayReverse - Return an array with elements in reverse order
// Original : https://www.php.net/manual/en/function.array-reverse.php
func ArrayReverse(v []interface{}) []interface{} {
		for ind, fnd:= 0, len(v); ind < fnd; ind, fnd = ind+1, fnd-1 {
			v[ind], v[fnd] = v[fnd], v[ind]
		}
		return v
}

// ArrayCountValues - Counts all the values of an array
// Original : https://www.php.net/manual/en/function.array-count-values.php
func ArrayCountValues(v []interface{}) map[interface{}]uint {

		cnt := make(map[interface{}]uint)
		for _,s:= range v {
			if c, ok := cnt[s]; ok {
				cnt[s] = c+1
			} else {
				cnt[s] = 1
			}
		}

		return cnt
}

// ArrayMerge - Merge one or more arrays
// Original : https://www.php.net/manual/en/function.array-merge.php
func ArrayMerge(v ...[]interface{}) []interface{} {

	fArray := make([]interface{},0)
	for _, m := range v {
		fArray = append(fArray, m...)
	}
	return fArray
}

// ArrayValues - Return all the values of an array
// Original : https://www.php.net/manual/en/function.array-values.php
func ArrayValues(v ...[]interface{}) (value []interface{}) {
		var values []interface{}

		for _, val := range v {
			cIn := InArray(val, v)
				if!cIn {
					values = append(values, val)
				}
		}

		return values
}

// ArrayChunk - Split an array into chunks
// Original : https://www.php.net/manual/en/function.array-chunk.php
func ArrayChunk(v ArraySlice, size int) ArraySlice {

	length := len(v)
	count := int(math.Ceil(float64(length) / float64(size)))
	ret := make(ArraySlice, count)
	for i := 0; i < count-1; i++ {
		ret[i] = v[i*size : (i+1)*size]
		fmt.Println(ret)
	}
	ret[count-1] = v[size*(count-1):]

	return ret
}

// ArrayFlip - Exchanges all keys with their associated values in an array
// Original : https://www.php.net/manual/en/function.array-flip.php
// ArrayFlip returns an array in flip order, i.e. keys from array become values and values from array become keys.
func ArrayFlip(v map[interface{}]interface{}) map[interface{}]interface{} {

	z := make(map[interface{}]interface{})
	for ind, fnd := range v {
		z[fnd] = ind
	}

	return z
}

// Count - Count all elements in an array, or something in an object
// Original : https://www.php.net/manual/en/function.count.php
func Count(v []interface{}) int {
	return len(v)
}