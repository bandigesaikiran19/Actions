// Golang program to swap adjacent elements of a
// one-dimensional array

package main

import "fmt"

func main() {
	var temp int = 0
	arr := [...]int{0, 1, 2, 3, 4, 5}

	fmt.Printf("Array elements: \n")
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d ", arr[i])
	}

	//Swap adjacent elements
	for i := 0; i <= 4; {
		temp = arr[i]
		arr[i] = arr[i+1]
		arr[i+1] = temp
		i = i + 2
	}

	fmt.Printf("\nArray elements after swapping adjacent elements: \n")
	for i := 0; i <= 5; i++ {
		fmt.Printf("%d ", arr[i])
	}
}
