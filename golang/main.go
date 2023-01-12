package main

import(
	"fmt"
)

//loop over 1,000,000,000 items
func main(){	
	// i := 0
	// for next := true; next; next = i < 1000000000 {		
	// 	i++
	// }
	// fmt.Println(i)

	//optimised code - for loop instead of a while loop - ChatGPT!
	n := 0
	for n = 0; n < 1000000000; n++ {
		//n++
	}
	fmt.Println(n)


}