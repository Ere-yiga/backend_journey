package main

import "fmt"

func main(){
	mp := map[string][]int{"a": {1, 2, 3}}
	
	//for a key value pair

	mp["b"] = []int{4, 5}
	fmt.Println(mp)

	//how to delete a key
	delete(mp, "b")
	fmt.Println(mp)

	//unlike append function in slices, here I don't need to assign or create a variable.

	//how to check if a key is present/absent

	/* value, ok := mp["b"]
	fmt.Println(value, ok)
 */
	value, ok := mp["a"]
	fmt.Println(value, ok)
}