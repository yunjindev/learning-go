// PAGE NUMBER: 88/374 -- FEB 4, 2026
// LOCK IN 

package main

import "fmt"

func addTo(base int, vals ...int) []int {
	returnOut := make([]int, 0, len(vals))
	for _, v := range vals {
		returnOut = append(returnOut, base+v)
	}
	return returnOut
}
func main() {
	fmt.Println(addTo(3,3,4,3))
}
