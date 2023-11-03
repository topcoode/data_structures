package array

import "fmt"

func Adding() {
	// Declare an array
	array := [5]int{1, 2, 3, 4, 5}

	// Define the element and index where you want to insert
	elementToInsert := 10
	insertIndex := 2

	// Shift elements to the right to make space for the new element
	for i := len(array) - 1; i > insertIndex; i-- {
		fmt.Println("looping--->", array)
		array[i] = array[i-1]
		fmt.Println("looping--->", array)
	}
	/*
	   looping---> [1 2 3 4 5]
	   looping---> [1 2 3 4 4]
	   looping---> [1 2 3 4 4]
	   looping---> [1 2 3 3 4]
	   [1 2 10 3 4]
	*/
	// Insert the element at the specified index
	array[insertIndex] = elementToInsert

	// Print the modified array
	fmt.Println(array)
}
