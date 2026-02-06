// PAGE NUMBER: 61/374 -- FEB 4, 2026
// LOCK IN 

package main

import "fmt"

// TODO find out what the type is in a struct
func main() {
	type firstPerson struct {
		name string
		age int
}

	f := firstPerson {
		name: "bob",
		age: 50,
	}

	var g struct {
		name string
		ages int
	}
	g = f
	fmt.Println(f == g)
}
