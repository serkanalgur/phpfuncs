package phpfuncs

import "time"

/***********************
* BASIC TIME FUNCTIONS *
************************/

/**
*	Similar function of time() in PHP.
* Original : https://www.php.net/manual/en/function.time.php
* func Time
* return current Unix timestamp
**/

// Time Function
func Time() int64{
	return time.Now().Unix()
}

/**
*	Similar function of time() in PHP.
* Original : https://www.php.net/manual/en/function.time.php
* func Now
* return current Unix timestamp
**/

// Now Function
func Now() int64{
	return time.Now().Unix()
}
