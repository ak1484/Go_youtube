package main

import (
	"fmt"
	// "os"
	// "strconv"
) 
func main() {
	fmt.Println("Maps")
	//Maps
	var mp map[string]int=map[string]int{
		"apple":5,
		"pear":7,
		"orange":9,
	}
	
	fmt.Println(mp)
	for key,value:=range mp{
		fmt.Println(key,"",value)
	}
	delete(mp,"apple")
	fmt.Println("\n",mp)
	val,ok:=mp["apple"]
	fmt.Println(val,ok)
}