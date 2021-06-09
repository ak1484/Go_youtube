package main

import (
	"fmt"
	// "os"
	// "strconv"
) 
type Student struct{
	name string
	grades []int
	age int
	
}
func (s Student) GAG()float64{
	sum:=0
	for _,v:=range s.grades{
		sum+=v
	}
	return float64(sum)/float64(len(s.grades))
}
func main() {
	fmt.Println("Struct Methods")
	// Struct Methods
    s1:=Student{"ankit",[]int{76,85,69},19}
	val:=s1.GAG()
	fmt.Println(val)

}