// PAGE NUMBER: 44/374 -- FEB 3, 2026
// LOCK IN 

package main

import ( "fmt" )

func main() {
	x := make([]int,0,8) // 0 ints, 5 cap
	x = append(x, 1,2,3,4)
	y:= x[:2] // 1 2
	z := x[2:] // 3 4
	fmt.Println(cap(x), cap(y), cap(z)) //5 5 3
	y = append(y, 30,40,50)
	x = append(x,60)
	z = append(z,70)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

}
