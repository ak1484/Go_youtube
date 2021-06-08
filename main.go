package main

import (
	"fmt"
	// "os"
	// "strconv"
)
func main() {
	fmt.Println("Chained Conditionals (AND, OR, NOT)")
	// || or, && and , ! not start from left to right evaluating
	val:=false && (false || true)|| ! (8>7)
	fmt.Printf("%t",val)
}