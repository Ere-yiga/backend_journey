package main

import "fmt"

type Shape interface{
	getPerimeter() uint
}

type Triangle struct{
	a uint
	b uint
	c uint
}

func (t Triangle) getPerimeter() uint{
	return t.a + t.b + t.c
}
/* 
func (t Triangle) getSides() []uint{
	return []uint(t.a + t.b + t.c)
} */

func main(){
	//you can use this or

	/* var s Shape = Triangle{2, 3, 5} */
	
	//...this
	
	s := Triangle{2, 3, 5}
	fmt.Println(s.getPerimeter())
}