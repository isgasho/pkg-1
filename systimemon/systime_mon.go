package systimemon

import (
	"time"
)

// StartMonitor will call systimeErrHandler if system time jump backward.
func StartMonitor(now func() time.Time, systimeErrHandler func()) {
	tick := time.NewTicker(100 * time.Millisecond)
	defer tick.Stop()
	for {
		last := now()
		select {
		case <-tick.C:
			if now().Sub(last) < 0 {
				systimeErrHandler()
			}
		}
	}
}
