package main

import (
	"fmt"
	// "os"
	// "strconv"
)

// func strchange(str *string){
// 	*str="changed"
// }
// func str1change(str1 string){
// 	str1="changed"
// }
func main() {
	fmt.Println("Pointers & Derefrence Operator (& and *)")
	//Pointers & Derefrence Operator (& and *)
	// x:=8
	// fmt.Println(&x)
	// x:=7
	// y:=&x
	// fmt.Println(x,y)
	// *y=9
	// fmt.Println(x)
	// x:="no change"
	// y:="no change"
	// strchange(&x)
	// str1change(y)
	// fmt.Println(x)
	// fmt.Println(y)
    // x:=9
	// y:=&x
	// fmt.Println(y)
	x:="no change"
	var y *string=&x
	fmt.Println(&y,"",y,"",*y)
}