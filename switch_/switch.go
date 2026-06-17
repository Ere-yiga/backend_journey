package main

import "fmt"

func main(){
	a := 5

	switch a{
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("default")
	} 

	// unlike JS, you can also rewrite your switch statemeent, in the sense that there'll be no need to put a variable, and in each case statement, you put in a condition

	b := 11
	
	switch {
		case b < 2:
			fmt.Println(b, "is less than two")

		case b == 2:
			fmt.Println(b, "is equal to two")

		default: 
			fmt.Println(b, "is greater than two")
	}

}

//In Go, there is practically no need to use a break, since Go automatically does it for you

