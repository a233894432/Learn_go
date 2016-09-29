/*
Project:Readhistory
Time:2016/9/29-15:11
Author:diogo
*/
// URLs provide a [uniform way to locate resources](http://adam.heroku.com/past/2010/3/30/urls_are_the_uniform_way_to_locate_resources/).
// Here's how to parse URLs in Go.

package main

import "fmt"
import "net"
import "net/url"

func main() {

	// We'll parse this example URL, which includes a
	// scheme, authentication info, host, port, path,
	// query params, and query fragment.
	s := "postgres://userd:passd@host.com:5432/path?k=v#f"


	// Parse the URL and ensure there are no errors.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Accessing the scheme is straightforward.
	fmt.Println(u.Scheme) // postgres

	// `User` contains all authentication info; call
	// `Username` and `Password` on this for individual
	// values.
	fmt.Println(u.User)            // userd:passd
	fmt.Println(u.User.Username()) // userd
	p, _ := u.User.Password()
	fmt.Println(p) // passd

	// The `Host` contains both the hostname and the port,
	// if present. Use `SplitHostPort` to extract them.
	fmt.Println(u.Host) // host.com:5432
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host) //  host.com
	fmt.Println(port) // 5432

	// Here we extract the `path` and the fragment after
	// the `#`.
	fmt.Println(u.Path)     // path
	fmt.Println(u.Fragment) // f

	// To get query params in a string of `k=v` format,
	// use `RawQuery`. You can also parse query params
	// into a map. The parsed query param maps are from
	// strings to slices of strings, so index into `[0]`
	// if you only want the first value.
	fmt.Println(u.RawQuery) // k=v
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)         // map[k:[v]]
	fmt.Println(m["k"][0]) // v
}
