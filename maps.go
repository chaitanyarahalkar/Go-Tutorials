package main 

import "fmt"


type Coordinates struct {
	Lat, Long float64
}


var m map[string]Coordinates
func main(){

	/* Map acts as a associative array in Go
		The make function is used to create maps
	*/
	m = make(map[string]Coordinates)
	m["Home"] = Coordinates{40.5061,-56.4141}

	fmt.Println(m)

	// Map literals- Alternate way of creating maps
	var m = map[string]Coordinates{
		"Bell Labs": {40.6843, -74.3996},
		"Google": { 37.4220, -122.0840},
	}
	fmt.Println(m)


	// Mutating maps

	// Insert 
	m["Office"] = Coordinates{11.3121,-12.3111}

	// Retrieve element

	var coordinates = m["Google"]
	fmt.Println("Coordinates of Google Inc: ",coordinates)

	// Delete element

	delete(m,"Bell Labs")

	// Check if element is present 

	var elem, ok = m["Microsoft"]

	/* Since element is not present in the map, the struct retuns
	    0,0 as the default coordinates and ok is false 
	*/
	fmt.Println(elem,ok)
}