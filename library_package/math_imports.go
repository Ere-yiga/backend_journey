// there are other ways you can import in Go


package main

import (
	"fmt"
	"math"

	//conversion to string
	"strconv"
)

func main(){
	//square root
	fmt.Println(math.Sqrt(9))

	//exponent
	fmt.Println(math.Pow(3, 3))

	a := "1234"
	c, err := strconv.Atoi(a)
	fmt.Println(c, err)

	//we also have parseint.
	//it comes with somethings though, the value, base and bits

	b := "234"
	d, err := strconv.ParseInt(b, 10, 0)
	fmt.Println(d)
}
//summary

//types that go in, should be out, we can use a type conversion to use operand