//Pointers allows us to view both the memory address location of where something exists, and to modify waht is held at that location.

//pointers reference to the address of the former

//pointers are being accessed through the ampersand symbol "&"

//you can dereference using the "*" ...it goes to the memory location of the former, then accesses the value.

package main

import "fmt"

/* func main(){
	x := 49
	y := &x

	*y = 100

	fmt.Println(x, *y)
} */



//pointers in function


/* 
func change(x *int){
	*x = 100
}

func main(){
	a := 10
	change(&a)

	fmt.Println(a)
} */



//modifying structs with pointer

/* type Book struct {
	id int
	title string
}

func (b *Book) setTitle(title string){
	b.title = title
}

func main(){
	b := Book{10, "Old"}
	b.setTitle("New")
	fmt.Println(b)
} */


//pointers to pointers

func fmrmain(){
	/* x := 0
	y := &x
	z := &y */

	/* fmt.Println(x, y, z) */
	/* fmt.Println(x, *y, *z) */
	/* fmt.Println(x, *y, **z) */
}

func test(pointerSlice *[]*int){
	values := *pointerSlice

	for _, value := range values{
		value
	}
}

func main(){
	a := 1
	b := 2
	values := &[]*int(&a, &b)
	fmt.Println()
}


//astericks as a type is a pointer

//asterick before variable, you're dereferencing

//ampersand before value, is referencing a value to give pointer

