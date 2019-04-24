package main
/* Go has only one looping construct - For 
The basic for loop has three components separated by semicolons:

the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration

Braces are always required for loops.


*/

import "fmt"



func main(){
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		sum += i
	}

	fmt.Println(sum)

	// The init and post statements are optional

	sum = 1
	for ;sum<=1000;{
		sum+=sum
	}
	fmt.Println(sum)


	// Semicolons can also be dropped 
	sum = 1
	for sum < 1000 {
		sum+=sum
	}
	fmt.Println(sum)



	// Infinite loops with a breaking condition
	sum = 2
	for{	
		if sum > 1000{
			break
		}
		sum*=sum
	}	
	fmt.Println(sum)
}