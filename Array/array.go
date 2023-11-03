package array

import "fmt"

func Array() {

	//Representation of Array
	//static array
	var StatiAarray = []int{65, 545, 67}
	fmt.Println(StatiAarray)
	//-------------------------------------------->
	//dynamic array
	var DynamicArray = [4]int{67, 56, 65}
	fmt.Println(DynamicArray)
	//------------------------------------------->

	//one dimensional array

	OnedimensionalArray := [4]string{"coding", "is", "good", "habit"}
	/*
		coding / 0
		is     / 1
		good   / 2
		habit  / 3
	*/
	for i := 0; i < 4; i++ {
		fmt.Println(OnedimensionalArray[i])
	}
	fmt.Println(OnedimensionalArray)
}
