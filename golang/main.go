package main

import(
	"fmt"
)

//loop over 1,000,000,000 items
func main(){
	const max_loop = 1000000000	
	i := 0
	for next := true; next; next = i < max_loop {		
		i++
	}
	fmt.Println("completed")
}