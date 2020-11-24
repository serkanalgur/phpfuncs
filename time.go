package phpfuncs

import "time"

/***********************
* BASIC TIME FUNCTIONS *
************************/

// Time - Similar function of time() in PHP.
// Original : https://www.php.net/manual/en/function.time.php
func Time() int64 {
	return time.Now().Unix()
}

// Now - Similar function of time() in PHP.
// Original : https://www.php.net/manual/en/function.time.php
func Now() int64 {
	return time.Now().Unix()
}

// Sleep - Delay in seconds
// Original : https://www.php.net/manual/en/function.sleep.php
func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}

// USleep - Delay in microseconds
// Original : https://www.php.net/manual/en/function.usleep.php
func USleep(t int64) {
	time.Sleep(time.Duration(t) * time.Microsecond)
}
