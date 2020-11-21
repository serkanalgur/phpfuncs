package phpfuncs

import "time"

// Sleep - Delay in seconds
func Sleep(t int64) {
	time.Sleep(time.Duration(t) * time.Second)
}

// USleep - Delay in microseconds
func USleep(t int64){
	time.Sleep(time.Duration(t) * time.Microsecond)
}