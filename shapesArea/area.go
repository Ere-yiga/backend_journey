package main

import "fmt"

type Shape interface{
	Area() float64
}

type Circle struct{
	pi float64
	r float64
}

type Rectangle struct{
	a float64
	b float64
}

func (c Circle) Area() float64{
	return c.pi * c.r * c.r
} 

func (d Rectangle) Area() float64{
	return d.a * d.b
} 

func main(){
	c := Circle{pi: 3.142857, r: 2}
	d := Rectangle{a: 3, b: 2}
	fmt.Println(c.Area())
	fmt.Println(d.Area())
/* 
	var s Shape
	s = c
	fmt.Println(s.Area())

	s = d
	fmt.Println(s.Area()) */
} 

//lessons this work taught me...

//know when to use uint and float. The former is for whole number, while the later is for decimal.

//know when do youse your curly bracess and normal bracket {} -- strct value, () -- calling functions or converting types

//te c Circle and co, function just like the way this, in JS, using c. is more like accessing via receiver variable. It doesn't have any special meaning. Can be myCircle...x...foo...nothing.


//for the instances, you can just type them out, buh using pi: makes it safer and more readable