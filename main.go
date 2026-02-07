// PAGE NUMBER: 73/374 -- FEB 4, 2026
// LOCK IN 

package main

import ( 
	"fmt"

)

func main() {
	samples := []string{"Hello", "apple_n!"}
	for x, sample := range samples {
		for i, r := range sample {
			fmt.Println(x,i,r, string(r))
		}
		fmt.Println()
	}
}

