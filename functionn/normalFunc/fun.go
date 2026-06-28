package main

import "fmt"

//In golang, you call out the data type

func div(num1 int, num2 int) int{
	return num1 / num2
}


//you can return multiple 
func many(one int, two int)(int, string){
	return one + two, "Hey there"
}

func main(){
	ans := div(9, 3)
	fmt.Println(ans)

	result, hei := many(1, 2)
	fmt.Println(result, hei)
}

//write out the name of the function, specify, write out parameters, return value
