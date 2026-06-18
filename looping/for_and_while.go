package main

import "fmt"

func main(){
	//for loop, same way you write in JS

	//initialization; condition; post
	for idx := 2; idx < 10; idx++{
		fmt.Println(idx)
	}

	//in Go, there is no while loop. Instead, the for loop is being used

	//you use a for  loop, then put in a single condition

	a := 1
	for a <= 10{
		fmt.Println("loop")
		a++
	}


	//in assessing strings, it is quite different here

	str := "hello world"
	fmt.Println(string(str[0]))

	//looping through a string is usually done using a range syntax

	//a lil exercise to know how many vowels are in a string

	vowels := 0

	for _, char := range "emmanuel"{
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			vowels++
		}
	}
	fmt.Print(vowels)
}