// PAGE NUMBER: 30/374 -- FEB 3, 2026
// LOCK IN 

package main

import ( "fmt" )

var x int64 = 10

var (
	IdKey = "id"
	NameKey = "name"
)
var help int64 = 20

func main() {
	var y = "hello"

	fmt.Println(x) // 10
	fmt.Println(y) // hello

	x = x + 1
	y = "blue"

	fmt.Println(x) // 11 but will fail cause of const
	fmt.Println(y) // blue but will fail
}


