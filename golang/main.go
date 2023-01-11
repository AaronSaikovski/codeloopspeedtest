package main

import(
	"fmt"
)

//loop over 1,000,000,000 items
func main(){	
	i := 0
	for next := true; next; next = i < 1000000000 {		
		i++
	}
	fmt.Println(i)
}