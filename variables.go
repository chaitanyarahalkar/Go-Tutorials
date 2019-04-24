/*
Datatypes in Go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte  alias for uint8

rune  alias for int32
      represents a Unicode code point

float32 float64

complex64 complex128


*/



package main

import "fmt"
import "math/cmplx"


/* Golang is statically as well as dynamically typed 
	The datatype is written at the end
	Booleans are assigned with false as the
    default value. Integers and floats with 0 & 0.0
    respectively. Strings have a default of ""

*/
var c, python, java bool


// Variables with initalizers 
var x,y int = 1,2

func main() {
	var i int

	var p,q,r,s = "Hey","Hello",1.1,false

	fmt.Println(i, c, python, java)

	fmt.Println(x,y)

	fmt.Println(p,q,r,s)


	/* Go supports short variable declarations. 
		The scope of these variables lie within the
		function in which they are declared.
	*/

	t := 10
	lx,yx := "abc","prq"
	fmt.Println(t)
	fmt.Println(lx,yx)


	// Factoring of variables

	var (
		ToBe bool = false
		MaxInt uint64 = 1<<64 - 1
		z complex128 = cmplx.Sqrt(-5 + 12i)
	)
	// %T gives the datatype of the variable and %v gives its value
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)


	// Type conversions

	var it int = 43
	f:= float64(it)
	fmt.Println(it,f)

	/* Constants can be characters,strings,boolean or numbers
		They cannot be declared using := syntax
	*/

	const num = 10
	const pi = 3.14
	const Truth = true 

	fmt.Println(num,pi,Truth)


	/* Numeric constants are high precision values
		An untyped constant. Constants can be declared 
		in factored format as well.
	*/

	const(
		Big  = 1 << 100
		Small = Big >> 99
	)
	fmt.Println(Big*0.1,Small)

}

