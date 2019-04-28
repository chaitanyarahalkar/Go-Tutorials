package main 

import "fmt"

func main(){

	/* A pointer holds the memory address of a value 
		A non initialized pointer in Go has the value 'nil'
		Go has no pointer arithmetic.

	*/

	var p *int

	i := 33
	p = &i

	fmt.Println(*p)

	// Dereferencing or indirecting 

	*p = 201

	fmt.Println(*p)


}