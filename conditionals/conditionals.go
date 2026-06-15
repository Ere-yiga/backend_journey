//condition, just like JS, evaluates true and falsed

package main

import "fmt"

func main(){
	/*
	<
	>
	<=
	==
	!=
	*/

	//logical

	//	||, &&, !

	x := 3
	if x < 3{
		fmt.Println("x is less than three")
	} else if x > 4{
		fmt.Println("false")
	}else{
		fmt.Println("X is equal to three")
	}

}