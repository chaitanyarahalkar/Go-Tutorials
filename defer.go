package main 


import "fmt"



func main(){

	/*
		A defer statement defers the execution of a function 
		until the surrounding function returns.

		The deferred call's arguments are evaluated immediately, 
		but the function call is not executed until the surrounding 
		function returns.

		If there are multiple defer calls, they are pushed onto a 
		stack. When the surrounding function returns a value,the 
		deferred calls are executed in LIFO order.

	*/

	defer fmt.Println("I will print after Hello World and the for loop!")

	fmt.Println("Hello World!")



	// Stacked defers

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}




}