package phpfuncs

import "time"

/***********************
* BASIC TIME FUNCTIONS *
************************/

// Time - Similar function of time() in PHP.
// Original : https://www.php.net/manual/en/function.time.php
// Returns the current time measured in the number of seconds since the Unix Epoch (January 1 1970 00:00:00 GMT).
func Time() int64 {
	return time.Now().Unix()
}

// Now - Similar function of time() in PHP.
// Original : https://www.php.net/manual/en/function.time.php
// Returns the current time measured in the number of seconds since the Unix Epoch (January 1 1970 00:00:00 GMT).
func Now() int64 {
	return time.Now().Unix()
}

// Sleep - Delay in seconds
// Original : https://www.php.net/manual/en/function.sleep.php
// Delays the program execution for the given number of seconds.
func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}

// USleep - Delay in microseconds
// Original : https://www.php.net/manual/en/function.usleep.php
// Delays program execution for the given number of microseconds.
func USleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}
