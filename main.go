package main

import (
	// "bufio"
	"fmt"
	// "os"
	// "strconv"
)
func main() {
	fmt.Println("Conditions & Boolean Expressions")
	//<,>,>=,<=,!=,== are all condition
	x:=5
	y:=6.4
	ans:=float64(x)!=y
	fmt.Printf("%t",ans)
	a:="tim"
	b:="tim"
	c:="Tim"
	fmt.Printf(" %t",a==b)
	fmt.Printf(" %t",a==c)
}