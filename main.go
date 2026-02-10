// PAGE NUMBER: 90/374 -- FEB 4, 2026
// LOCK IN 

package main

import ( 
	"fmt"
	"errors"
	"os"
)

func divAndRemainder(num int, denom int) (int,int,error) {
	if denom == 0 {
		return 0,0, errors.New("cannot divide by 0")
	}
	return num / denom, num % denom, nil
}

func main() {
	result, remainder, err := divAndRemainder(_,_)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(result, remainder)
}
