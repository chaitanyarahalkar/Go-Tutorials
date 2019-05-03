package main 



import "fmt"

var pow = []int{1,2,3,4,5,6,7,8,9,10}

func main(){


	fmt.Println("Range based iteration")
	// Range starts from 0 till the end of the iterable.
	for i,v := range pow {
		fmt.Println(i,v)
	}

	fmt.Println("Range: Using '_' as the iterator variable")
	// Useless variables can be replaced with '_'
	for _,v := range pow{
		fmt.Println(v)
	}

	fmt.Println("Range: Without the iterable")
	// Useless variables can also be completely eliminated
	for v := range pow{
		fmt.Println(v)
	}


}