// PAGE NUMBER: 41/374 -- FEB 3, 2026
// LOCK IN 

package main

import ( "fmt" )

func main() {
	var x []int
	var y = []int{20, 30}
	x= append(x,y...)
	fmt.Println(x)
}
