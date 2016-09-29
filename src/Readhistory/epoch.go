/*
Project:Readhistory
Time:2016/9/29-14:48
Author:diogo
*/
// A common requirement in programs is getting the number
// of seconds, milliseconds, or nanoseconds since the
// [Unix epoch](http://en.wikipedia.org/wiki/Unix_time).
// Here's how to do it in Go.

package main

import "fmt"
import "time"

func main() {

	// Use `time.Now` with `Unix` or `UnixNano` to get
	// elapsed time since the Unix epoch in seconds or
	// nanoseconds, respectively.
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now) //2016-09-29 14:48:59.7678239 +0800 CST

	// Note that there is no `UnixMillis`, so to get the
	// milliseconds since epoch you'll need to manually
	// divide from nanoseconds.
	millis := nanos / 1000000
	fmt.Println(secs)   // 1475131739
	fmt.Println(millis) // 1475131739767
	fmt.Println(nanos)  // 1475131739767823900

	// You can also convert integer seconds or nanoseconds
	// since the epoch into the corresponding `time`.
	fmt.Println(time.Unix(secs, 0))  // 2016-09-29 14:48:59 +0800 CST
	fmt.Println(time.Unix(0, nanos)) // 2016-09-29 14:48:59.7678239 +0800 CST
}
