package main

import (
	"fmt"
	// "os"
	// "strconv"
) 
func main() {
	fmt.Println("Range & Slice/Array Examples")
	//Range & Slice/Array Examples
	s:=[]int{1,2,3,5,6,5,3}
	// for i:=0;i<len(s);i++{
	// 	fmt.Println(s[i])
	// }
	for i,ele:=range s{
		for _,ele2:=range s[i+1:]{
			if ele2==ele{
				fmt.Println(ele)
			}
		}
	}
}