package main

import (
	"fmt"
	// "os"
	// "strconv"
) 
func main() {
	fmt.Println("Slices")
	//Slices
    // var arr[5]int=[5]int{1,2,3,5,6}
	// var S[]int=arr[1:3]
	// fmt.Println(S)
	// fmt.Println(cap(S))
	// fmt.Println(S[0:1])
	// var s []int=[]int{1,2,3,5,7}
	// fmt.Println(cap(s))
	// x:=append(s,10)
	// fmt.Println(x)
	a:=make([]int,5)
	fmt.Println(a)
	fmt.Printf("%T",a)
}