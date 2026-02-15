// PAGE NUMBER: 130/374 -- FEB 13, 2026
// LOCK IN 

package main

import ( 
	"fmt"
)
type Person struct {
	FirstName string
	LastName string
	Age int32
	Count int64
}

	func (p *Person) string() string {
		p.Count++
		return fmt.Sprintf("%s %s, age: %d, Count: %d", p.FirstName, p.LastName, p.Age, p.Count)
	}



func main() {
	douglas := Person {
		FirstName: "Douglas",
		LastName: "White",
		Age: 33,
		Count: 12,
	}
	output := douglas.string()
	fmt.Println(output, douglas.Count)	
}
