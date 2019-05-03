package main 


import (
	"fmt"
	"strings"
	)



func main(){


	// Syntax var name[count]Type 

	var a[2]string 
	a[0] = "Hello"
	a[1] = "世界"

	// Arrays cannot be resized in Go
	fmt.Println(a)


	// A different approach of declaring arrays 
	primes := [6]int{2, 3, 5, 7, 11, 13}


	fmt.Println(primes)

	// Array slicing in Go
	var s []int = primes[1:4]

	fmt.Println(s)


	/* Slices are references to arrays and changing the 
		element of slice modifies the corresponding elements
		of the underlying array
	*/

	s[0] = 17
	fmt.Println(s,primes)

	/*
		When slicing, you may omit the high or low bounds to 
		use their defaults instead.
		The default is zero for the low bound and the length 
		of the slice for the high bound.

		All the declarations given below produce the same result
	*/
	
	fmt.Println(primes[:])
	fmt.Println(primes[0:])
	fmt.Println(primes[:6])
	fmt.Println(primes[0:6])


	// Slice literals 


	v := []int{1,3,5,7,9,11}
	fmt.Println(v)

	/* Length of a slice is the number of elements in the slice
		Capacity of a slice is the number of elements in the underlying array
	*/
	SizeOfSlice(s)


	// Nil slice 

	var n []int

	if n == nil{
		fmt.Printf("The slice: %v is nil!\n",n)
	}

	/* 	The built in function make allows to create slices in Go
		Dynamically sized arrays can be built using make
		Make creates an array filled with zeroes

		make([]type,length,capacity)
	*/

	d := make([]int,10)
	SizeOfSlice(d)
	e := d[:5]
	SizeOfSlice(e) 

	f := make([]int,0,5)
	SizeOfSlice(f)


	// A two dimensional string matrix 
	grid := [][]string{
		[]string{"1", "2", "3"},
		[]string{"4", "5", "6"},
		[]string{"7", "8", "9"},
	}
	for i := 0; i < len(grid); i++ {
		fmt.Println(strings.Join(grid[i]," "))
	}

	// Built in append function 

	d = append(d,1,2,3,4,5)

	fmt.Println(d)
	
}

func SizeOfSlice(a []int){
	fmt.Printf("For slice: %v length of Slice=%d and capacity of underlying array=%d\n",a,len(a),cap(a))
}