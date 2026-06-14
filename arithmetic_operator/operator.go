package main

import "fmt"

func main(){
	a := 7
	b := 1
	c := a - b

	fmt.Println(c)

	//conversion using different types
	//it is done from smaller number to the bigger one, that's why we try to convert from uint8
	
	x := uint8(7)
	y := 1000
	z := int(x) + y
	fmt.Println(z)

	//division
	//you can used either 32 or 64, buh the latter is recommended
	
	p := 7
	q := 1000
	r := float64(q) / float64(p)
	fmt.Println(r)

	//concatenation
	//adding a string to an int, you convert using a fmt.sprint
	m := "hello"
	n := 5
	o := m + fmt.Sprint(n)
	fmt.Println(o) 
}