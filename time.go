package phpfuncs

import "time"

/***********************
* BASIC TIME FUNCTIONS *
************************/

// Time Function
func Time() int64{
	return time.Now().Unix()
}

// Now Function
func Now() int64{
	return time.Now().Unix()
}