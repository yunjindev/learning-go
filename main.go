// PAGE NUMBER: 101/374 -- FEB 10, 2026
// LOCK IN 

package main

import ( 
	"fmt"
)
func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}

func main() {
	twoBase := makeMult(5)
	threeBase := makeMult(3)
	for i := 0; i < 3; i++ {
		fmt.Println(twoBase(i), threeBase(i))
	}
}
