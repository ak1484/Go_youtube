package main

import (
	"fmt"
	// "os"
	// "strconv"
)
func main() {
	fmt.Println("Arrays")
	//Arrays
    var arr [3]int
	//default 0 is initilized
	arr[0]=100
	fmt.Println(arr)
	arrr:=[5]int{2,4,5,8,9}
	for i:=0;i<len(arrr);i++{
		fmt.Printf("%d\n",arrr[i])
	}
}