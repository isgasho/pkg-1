// Copyright 2016 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package systimemon

import (
	"time"
)

// StartMonitor will call systimeErrHandler if system time jump backward.
func StartMonitor(now func() time.Time, systimeErrHandler func() bool) {
	tick := time.NewTicker(100 * time.Millisecond)
	defer tick.Stop()
	for {
		last := now()
		select {
		case <-tick.C:
			if now().Sub(last) < 0 {
				if !systimeErrHandler() {
					tick.Stop()
					break
				}
			}
		}
	}
}
