package main 
 

 import "fmt"



 func main(){
 	// A structure is a collection of fields

 	type Coordinates struct{
 		x int
 		y int 
 	}
 	v := Coordinates{1,2}


 	// Access variables with dot 

 	v.x = 10
 	fmt.Println(v.x,v.y)


	// Pointer to struct 

	p := &v
	p.x = 100
	p.y = 150 

	fmt.Println(p)

	// Assignment of variables in structs
	var (
	v1 = Coordinates{1, 2}  // has type Coordinates
	v2 = Coordinates{x: 1}  // Y:0 is implicit
	v3 = Coordinates{}      // X:0 and Y:0
	p1  = &Coordinates{1, 2} // has type *Coordinates
	)

	fmt.Println(v1, p1, v2, v3)

 }