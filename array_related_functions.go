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