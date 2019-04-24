package main



import "fmt"
import "math"
import "runtime"
import "time"

// Shorthand notation of if statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}



func main(){

	/* Braces around the if statement are optional but
	 	brackets are mandatory. 
	 */
	var x int = 100
	if x < 1000 {
		fmt.Println("x is less than 1000")
	}


	
	fmt.Println(pow(3, 2, 10),pow(3, 3, 20))


	// If-Else block 

	if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
    }



    /* Switch case statements 
    	Go switch statements do not require 
    	break. They are auto included in Go. 
    	The case variables can be strings,integers,characters
    	and boolean.The case variables can also include 
    	functions provided that they have a return type.

    	Swtich case statements are evaluated from top to bottom
    */

    switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Go is running on OS X.")
	case "linux":
		fmt.Println("Go is running on Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("Go is running on %s.\n", os)
	}

	/* Switch may not have a condition variable inside it. Switch 
		without a condition is same as switch with true inside it.
	*/
	t := time.Now()

	switch{
		case t.Hour() < 12: 
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon!") 

		default:
			fmt.Println("Good evening")
	}

}