// PAGE NUMBER: 71/374 -- FEB 4, 2026
// LOCK IN 

package main

import ( 
	"fmt"

)

func main() {
	for i := 1; i < 101; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz ", i)
		}
		if i % 5 == 0 {
			fmt.Println("Buzz", i)
		}
		if i % 3 == 0 {
			fmt.Println("Fizz", i)
			}
		fmt.Println(i)
		}	

	}

