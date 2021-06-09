package main

import (
	"fmt"
	"math"
	// "os"
	// "strconv"
) 
type Rect struct{
	l float64
	w float64
}
type Circle struct{
	r float64
}
func (r Rect)Area()float64{
	return r.l*r.w
}
func (c Circle)Area()float64{
	return math.Pi*c.r*c.r
}
type shape interface{
	Area()float64
}
func main() {
	fmt.Println("Interfaces")
	//Interfaces
    r:=Rect{3,4}
	c:=Circle{3}
    s:=[]shape{r,c}
	for _,shape:=range s{
		fmt.Println(shape.Area())
	}
	fmt.Println(r.Area())
}