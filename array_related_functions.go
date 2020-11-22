package phpfuncs

// Array - Create an array
func Array(v... interface{}) []interface{} {
	return v
}

// ArrayKeys - Return array keys. Probably return keys as string
func ArrayKeys(v map[string]interface{}) []string {
	var aRdata []string

	for s := range v {
		aRdata = append(aRdata, s)
	}

	return aRdata
}

// ArrayPush - Add an element to end of given array
func ArrayPush(v *[]interface{}, data ...interface{}) {
		*v = append(*v, data...)
}

// ArrayReverse - Reverse an array
func ArrayReverse(v []interface{}) []interface{} {
		for ind, fnd:= 0, len(v); ind < fnd; ind, fnd = ind+1, fnd-1 {
			v[ind], v[fnd] = v[fnd], v[ind]
		}
		return v
}

// ArrayCountValues - Count an arrays values
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

// Count - Count array values
func Count(v []interface{}) int {
	return len(v)
}