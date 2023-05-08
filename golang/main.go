package main

import (
	"fmt"
)

//loop over 1,000,000,000 items
func main(){
	var x int=0
	for n:=0; n<1000000000; n++ {
		x++
	}
	fmt.Println(x)
}