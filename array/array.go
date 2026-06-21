package main

import "fmt"

func main(){
	var arr [2]int
	fmt.Println(arr)

	aaa := [2]int{31, 22}
	fmt.Println(aaa)

	//nested array
	num := [2][2]int{{1,2}, {3,4}}

	
	//properties of array
	//accessing items
	num[0] = [2]int{11, 12}

	fmt.Println(len(num))

	//looping through
	//you can only loop through a structure with more than one layer
	for _, a := range num{
		for _, b := range a{
			fmt.Print(b)
		}
	}

}