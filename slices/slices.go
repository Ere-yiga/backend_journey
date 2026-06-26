package main

import "fmt"

func main(){

	//pointer
	//length
	//capacity

	arr := [5]int{0, 1, 2, 3, 4}
	sli := arr[:3]	
	fmt.Println(sli)

	//using without creating an underline array

	sl := []string{"Hello", "world"}

	for x := 0; x < 3; x++{
		//reassign sl

		sl = append(sl, "tim")
	}
	fmt.Println(sl)

	//another way to create/make a slice
	sly := []string{"hello", "hi", "evening"}

	for _, value := range sly{
		fmt.Println(value)
	}
}
