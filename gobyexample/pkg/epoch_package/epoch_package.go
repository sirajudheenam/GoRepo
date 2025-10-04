package epochpackage

// A common requirement in programs is getting the number of seconds, milliseconds, or nanoseconds since the Unix epoch.
// Here’s how to do it in Go.

import (
	"fmt"
	"time"
)

func Run() {
	fmt.Println("\nEpoch: Time since 1970..")

	// Use time.Now with Unix, UnixMilli or UnixNano to get elapsed time since the Unix epoch in seconds, milliseconds or nanoseconds, respectively.

	now := time.Now()
	fmt.Println(now)

	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// You can also convert integer seconds or nanoseconds since the epoch into the corresponding time.

	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}

// $ go run main.go
// Epoch: Time since 1970..
// 2025-10-04 15:21:09.445633 +0200 CEST m=+0.000138918
// 1759584069
// 1759584069445
// 1759584069445633000
// 2025-10-04 15:21:09 +0200 CEST
// 2025-10-04 15:21:09.445633 +0200 CEST
