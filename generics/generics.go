package main

import "fmt"

//generics is simply a way we can have flexible types in statically typed language.

//Let's say you have an add function for an int, and you want to create other functions.

/* func Sub(a int, b int) int{
	return a - b
}

func Sub2(c int, d int) int{
	return d - c
}

func main(){
	fmt.Println(Sub(4, 3))
	fmt.Println(Sub2(3, 20))
} */

//Instead of creating those other functions, you can simply use a generic.

/* func add [T int | float64 | uint] (a T, b T) T{
	sum := a + b
	return sum
}

func main(){
	value1 := add(1, 9)
	value := add(3.3, 2.7)
	value2 := add(uint(3), uint(2))
	
	fmt.Print(value1, value, value2)
} */

//let's say we what to use multiple generic variables for like when we want to view a map.

/* func getValues[K comparable, V any](mp map[K]V) []V{

	values := []V{}

	for _, value := range mp{
		values = append(values, value)
	}
	return values	
} 

func main(){
	mp := map[string]int{
		"a": 100, "b": 80, "c": 60, "d": 40, "e": 20,
	}
	values := getValues(mp)
	fmt.Println(values)
} */

//common use cases for generics. Remember when we used the unions "|", that isn't the best case...

//just create an interface

/* type Number interface{
	int | float64 | uint
}

func add [T Number] (a T, b T) T{
	sum := a + b
	return sum
}

func main(){
	value1 := add(1, 9)
	value := add(3.3, 2.7)
	value2 := add(uint(3), uint(2))
	
	fmt.Print(value1, value, value2)
} */ 

//creating generic types

type GenericsSlices[T any] []T

func (g GenericsSlices[T]) Print() {
	for _, value := range g{
		fmt.Println(value)
	}
}

type GenericStruct[T any, K any] struct{
	values T
}

func main(){
	g := GenericsSlices[int]{1, 2, 3}
	g.Print()

	st := GenericStruct[string, int]{values: "string"}
	fmt.Println(st)
}
