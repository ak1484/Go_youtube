package main

import (
	"fmt"
	// "os"
	// "strconv"
)
func main() {
	fmt.Println("Switch Statement")
	//switch statement
    ans:=10
	for i:=1;i<ans;i++{
		rem:=ans%i
		switch rem{
		case 1:{
           fmt.Printf("%d",rem)
		}
	    case 2:{
		   fmt.Printf("%d",rem)
	    }
	    default:{
			fmt.Printf("%d",rem)
		}
		}
		fmt.Printf("\n")
	}
	
}