//Error here is quite different form JS and dynamica types languages.

//The errors in Go, will be caught at compile time,i.e, my codes won't run unless I fix majority of these issues.

//there is also run time error, those we didn't catch at compile time e.g panic -- dividing by error.


package main

import (
	"fmt"
	"errors"
)

//firstly, using a keyword defer. This shows that a function should happen after another has

//secondly, another keyword panic.This one helps display an error mesage. N/B after the panic lines, all other lines won't run

//another one is using the recover keyword. It is embedded inside the defferedFunc


/* func defferedFunc(){
	fmt.Print("defer")
	r := recover()
	fmt.Println(r)
}

func main(){
	defer defferedFunc()
	//panic("run wont run")

	fmt.Println("run")
} */

//check out runtime vs compile time errors

//If you notice above the codes are a bit too clungy, and in Go, we don;t ususally deal with run time errors.



//so rather than panicking, you just call the function

func divide(a int, b int) (int, error){
	if b == 0{
		return 0, errors.New("division by 0")
	}
	return a / b, nil
}

func main(){
	result, err := divide(4, 0)

	if err != nil{
		fmt.Println("error occured,", err)
	}else{
		fmt.Println(result)
	}
}