package main 

import "fmt"


/* Syntax of a function 

func func_name(variable datatype,...) return type {
	function body

	return value
}

*/
func add(x int,y int) int {
	return x+y
}

// Shorthand form can be used as shown below

func sub(x,y int) int {
	return x-y
}

// Returning multiple values from a function with parenthesis

func swap(x,y string)(string,string){
	return y,x
}


/* Named return values - Naked return
	Should be used only with short functions
	They harm the readability of longer functions
*/
func split(sum int)(x,y int){
	x = sum * 4/9
	y = sum - x
	return 
}

// Functions as values

func wrapper(fn func(int,int) int) int {
	return fn(21,31)
}
func main(){
	fmt.Println(add(5,6))


	fmt.Println(sub(6,4))

	
	a,b := swap("hello","world")
	fmt.Println(a,b)

	fmt.Println(split(17))


	fmt.Println("Wrapper function: Passed add() as an argument")
	fmt.Println(wrapper(add))

	fmt.Println(wrapper(sub))
}