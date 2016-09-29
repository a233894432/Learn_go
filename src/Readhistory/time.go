/*
Project:Readhistory
Time:2016/9/29-14:40
Author:diogo
*/
// Go's offers extensive support for times and durations;
// here are some examples.

package main

import "fmt"
import "time"

func main() {
	p := fmt.Println

	// We'll start by getting the current time.
	now := time.Now()
	p(now)

	// You can build a `time` struct by providing the
	// year, month, day, etc. Times are always associated
	// with a `Location`, i.e. time zone.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// You can extract the various components of the time
	// value as expected.
	p(then.Year())       // 2009
	p(then.Month())      //11
	p(then.Day())        // 17
	p(then.Hour())       //20
	p(then.Minute())     // 34
	p(then.Second())     //58
	p(then.Nanosecond()) // 651387237
	p(then.Location())   // UTC

	// The Monday-Sunday `Weekday` is also available.
	p(then.Weekday()) //Tuesday

	// These methods compare two times, testing if the
	// first occurs before, after, or at the same time
	// as the second, respectively.
	p(then.Before(now)) // true
	p(then.After(now))  // false
	p(then.Equal(now))  // false

	// The `Sub` methods returns a `Duration` representing
	// the interval between two times.
	diff := now.Sub(then)
	p(diff)   // 60178h 8m 22.347338263s

	// We can compute the length of the duration in
	// various units.
	p(diff.Hours()) //60178.139540927295
	p(diff.Minutes()) // 3.610688372455638e+06
	p(diff.Seconds()) // 2.1664130234733826e+08
	p(diff.Nanoseconds()) //216641302347338263

	// You can use `Add` to advance a time by a given
	// duration, or with a `-` to move backwards by a
	// duration.
	p(then.Add(diff))  //2016-09-29 06:43:20.9987255 +0000 UTC
	p(then.Add(-diff)) //2003-01-06 10:26:36.304048974 +0000 UTC
}
