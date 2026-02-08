// PAGE NUMBER: 76/374 -- FEB 4, 2026
// LOCK IN 

package main

import ( 
	"fmt"

)

func main() {
	samples := []string{"Hello", "apple_n!"}
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i,r, string(r))
		}
		fmt.Println()
	}
}

