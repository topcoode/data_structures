package array

import "fmt"

func Deletion_Array() {

	Array2 := []int{56, 82, 85, 72, 78}
	index := 2 //85->60
	AfterDeletion := append(Array2[:index], Array2[index+1:]...)
	/*
	   This line creates a new array newArray by
	   appending two slices of the original array.
	   The first slice is from the beginning of the array
	   up to the index of the element to delete (originalArray[:index]).

	   The second slice is from the element after the
	   index of the element to delete up to the end of
	   the array (originalArray[index+1:]).
	   The ... at the end of the line is used to unpack the slices into the append() function.
	*/
	fmt.Println("Array2[:index]--->", Array2[:index])     //Array2[:index]---> [56 82]
	fmt.Println("Array2[index+1:]--->", Array2[index+1:]) //Array2[index+1:]---> [78 78]
	fmt.Println(AfterDeletion)
}
