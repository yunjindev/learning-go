// PAGE NUMBER: 48/374 -- FEB 3, 2026
// LOCK IN 

package main

import ( "fmt" )

func main() {
	x := []int{1,2,3,4}
	y := make([]int,4)
	num := copy(y[:3],x[2:])
	fmt.Println(y,num)
}
