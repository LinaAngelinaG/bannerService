package schedule

import "time"

func Repeat(fn func() time.Duration) {
	go func() {
		for {
			time.Sleep(fn())
		}
	}()
}
