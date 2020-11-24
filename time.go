package phpfuncs

import "time"

/***********************
* BASIC TIME FUNCTIONS *
************************/

/**
* Time
* Similar function of time() in PHP.
* Original : https://www.php.net/manual/en/function.time.php
* func Time
* return current Unix timestamp
**/
func Time() int64 {
	return time.Now().Unix()
}

/**
* Now
* Similar function of time() in PHP.
* Original : https://www.php.net/manual/en/function.time.php
* func Now
* return current Unix timestamp
**/
func Now() int64 {
	return time.Now().Unix()
}

// Sleep - Delay in seconds
func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}

// USleep - Delay in microseconds
func USleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}
