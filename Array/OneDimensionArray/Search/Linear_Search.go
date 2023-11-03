package array

import "fmt"

func Linear_Array() bool {
	Array_Data := []int{1, 2, 3, 4, 5, 6, 7}
	Valuedata := 6
	for key, value := range Array_Data {
		//key - value
		fmt.Printf("key %v , value %v \n", key, value)
		if value == Valuedata {
			fmt.Println("\n", Array_Data[Valuedata])
			return true
		}
	}
	return false
}
func Linear_Main() {
	data := Linear_Array()
	fmt.Println("data:", data)
}
