// PAGE NUMBER: 111/374 -- FEB 12, 2026
// LOCK IN 

package main

import ( 
	"fmt"
)
type Person struct {
	FirstName string
	LastName string
	Age int32
}



func main() {
	douglas := Person {
		FirstName: "Douglas",
		LastName: "White",
		Age: 33,
	}
	fmt.Printf("This is %s %s, he is %d years old\n", douglas.FirstName, douglas.LastName, douglas.Age) 
}
