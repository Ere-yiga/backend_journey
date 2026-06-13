package main

import "fmt"

func main(){
	//this works just like console.log in JS
	fmt.Println("Hey")

	//we also have Printf, which means print format

	a := true

	// %T tells that the data type shpuld be printed/known 

	fmt.Printf("%T\n", a)


	//the output here is gonna be "bool"

	// %v tells value
	
	b := false

	fmt.Printf("the vale of b is: %v\n", b)
	//the output will be false

	
	//to get a percent sign, we do it twice

	fmt.Printf("23%%")

	//we also have Sprintf. This is when we want to pass in, then print later in a function
}

// there are several others
// %b, for binary, %e for scientific notation, %s for string value

//for rounding up, we have another one
//let's say we want to round up 7.822 to I decimal place, we use "%.1f. We'll get 7.8 as output"

//in summary, printf is for special formatting