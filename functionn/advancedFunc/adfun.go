package main

import "fmt"

//passing a fiuinction as a parameter to another function

func callFunc(callab func (int) int) int{
	return callab(10)
}

func doubleNum(num int) int{
	return 2 * num
}

func tripleNum(num int) int{
	return 3 * num
}

func main(){

	fmt.Println()
}