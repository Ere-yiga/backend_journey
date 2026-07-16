package main

import "fmt"

//passing a function as a parameter to another function

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
	value := callFunc(tripleNum)
	fmt.Println(value)
	
	//anonymous functions - they are also 	called nameless functions

	anonym := callFunc(func (x int) int{
		return x + 12
	})
	fmt.Println(anonym)
}

//look into named return value
//